import React, {useState, useCallback} from 'react';
import {gql, useQuery} from '@apollo/client';
import {View, StyleSheet, Alert} from 'react-native';
import Geolocation from '@react-native-community/geolocation';
import {geoFire, firebaseRef, distanceBetween} from 'services/firebase';
import StoreItem from './StoreItem';
import {StoreInfoUserQuery} from './__generated__/StoreInfoUserQuery';
import {StoresQuery, StoresQuery_stores} from './__generated__/StoresQuery';
import {Text, Button} from 'react-native-elements';
import ListWithError from 'components/ListWithError';
import {secToHour} from 'utils';
import {useAppDispatch, useAppSelector, setPost, increment} from 'services/store';
import {useNavigation, useFocusEffect} from '@react-navigation/native';
import {ScreenNames} from 'services/navigation';

const STORES_QUERY = gql`
  query StoresQuery {
    stores {
      ...store
      latitude
      longitude
    }
  }
  ${StoreItem.fragments.data}
`;

const STORE_INFO_USER_QUERY = gql`
  query StoreInfoUserQuery {
    currentUser {
      id
    }
  }
`;

const StoreInfo = () => {
  const [distances, setDistances] = useState({} as {[storeId in number]: number});
  const [waitingIntervalId, setWaitingIntervalId] = useState<NodeJS.Timer>();
  const [collapsed, setCollapsed] = useState(true);
  const navigation = useNavigation();
  const {loading, error, data, refetch} = useQuery<StoresQuery>(STORES_QUERY);
  const {data: currentUserData} = useQuery<StoreInfoUserQuery>(STORE_INFO_USER_QUERY);
  const post = useAppSelector(state => state.post);
  const dispatch = useAppDispatch();

  const updateDistances = useCallback((coords: {latitude: number; longitude: number}, stores: StoresQuery_stores[]) => {
    let newDistances = {} as {[storeId in number]: number};
    stores.forEach(store => {
      const distance = distanceBetween([coords.latitude, coords.longitude], [store.latitude, store.longitude]);
      newDistances = {...newDistances, [store.id]: distance};
    });
    setDistances(newDistances);
  }, []);

  useFocusEffect(
    useCallback(() => {
      if (!currentUserData || !data) {
        return;
      }
      dispatch(setPost({key: 'userId', value: currentUserData.currentUser.id}));
      const currentUserId = currentUserData.currentUser.id.toString();

      const watchId = Geolocation.watchPosition(
        position => {
          console.log(position);
          // This is just for an individual display. Leave the whole process of real-time suggesting to geofire.
          // This way you can easily integrate location data of other users. All comes in real-time.
          // @TODO can't we just directly get the distance data from a geofire ref?
          updateDistances(position.coords, data.stores);
          // Set current location to firebase realtime database.
          geoFire.set(currentUserId, [position.coords.latitude, position.coords.longitude]).then(() => {
            // When the user disconnects from Firebase (e.g. closes the app, exits the browser),
            // remove their GeoFire entry
            firebaseRef.child(currentUserId).onDisconnect().remove();
          });
        },
        err => Alert.alert('Error', JSON.stringify(err)),
        {enableHighAccuracy: true, timeout: 10000, maximumAge: 1000},
      );

      const cancels = [] as (() => void)[];
      data.stores.forEach(store => {
        const query = geoFire.query({
          center: [store.latitude, store.longitude],
          radius: 0.1,
        });
        query.on('key_entered', (key: string) => {
          console.log('someone entered');
          if (key === currentUserId && post.store.id === 0) {
            dispatch(setPost({key: 'store', value: store}));
          }
        });
        query.on('key_exited', (key: string) => {
          console.log('someone exited');
          if (key === currentUserId && post.store.id === store.id) {
            dispatch(setPost({key: 'store', value: {id: 0, name: ''}}));
          }
        });
        cancels.push(query.cancel.bind(query));
      });

      return () => {
        Geolocation.clearWatch(watchId);
        cancels.forEach(cancel => cancel());
        waitingIntervalId && clearInterval(waitingIntervalId);
      };
    }, [currentUserData, data, updateDistances, dispatch, post.store.id, waitingIntervalId]),
  );

  useFocusEffect(
    useCallback(() => {
      const intervalId = setInterval(() => {
        refetch();
      }, 60000);
      return () => clearInterval(intervalId);
      // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []),
  );

  const lineUp = () => {
    const intervalId = setInterval(() => {
      dispatch(increment('waitingFor'));
    }, 1000);
    setWaitingIntervalId(intervalId);
  };

  const getOut = useCallback(() => {
    dispatch(setPost({key: 'waitingFor', value: 0}));
    clearInterval(waitingIntervalId as NodeJS.Timer);
  }, [dispatch, waitingIntervalId]);

  if (post.waitingFor > 0) {
    return (
      <View style={styles.store}>
        <Text style={styles.storeName}>{post.store.name}</Text>
        <Text style={styles.waiting}>
          待機中...<Text style={styles.waitingFor}>{secToHour(post.waitingFor)}</Text>
        </Text>
        <View style={styles.actionGroup}>
          <View style={styles.invisible} />
          <Button title="着丼" onPress={() => navigation.navigate(ScreenNames.CreateImages)} />
          <Text style={styles.cancelSuggest} onPress={getOut}>
            キャンセル
          </Text>
        </View>
      </View>
    );
  } else if (post.store.id > 0) {
    return (
      <View style={styles.store}>
        <Text>
          今<Text style={styles.storeName}>{post.store.name}</Text>の近くにいますか？
        </Text>
        <View style={styles.actionGroup}>
          <View style={styles.invisible} />
          <Button title="並ぶ" onPress={lineUp} />
          <Text
            style={styles.cancelSuggest}
            onPress={() => dispatch(setPost({key: 'store', value: {id: 0, name: ''}}))}>
            いいえ
          </Text>
        </View>
      </View>
    );
  } else {
    return (
      <View style={styles.store}>
        <View style={styles.row}>
          <Text style={styles.name}>店名</Text>
          <Text style={styles.hour}>閉店まで</Text>
          <Text style={styles.distance}>お店まで</Text>
        </View>
        <View style={{height: collapsed ? 130 : 700}}>
          <ListWithError loading={loading} error={error?.message} objects={data?.stores}>
            {store => (
              <StoreItem
                data={store}
                distance={distances[store.id]}
                setCurrentStore={() => dispatch(setPost({key: 'store', value: store}))}
              />
            )}
          </ListWithError>
        </View>
        <View style={styles.collapseWrapper}>
          {collapsed ? (
            <Text style={styles.collapse} onPress={() => setCollapsed(false)}>
              全て表示
            </Text>
          ) : (
            <Text style={styles.collapse} onPress={() => setCollapsed(true)}>
              たたむ
            </Text>
          )}
        </View>
      </View>
    );
  }
};

const styles = StyleSheet.create({
  store: {
    padding: 20,
    borderRadius: 20,
    marginHorizontal: 25,
    marginTop: 10,
    marginBottom: 20,
    backgroundColor: '#E0E5E3',
  },
  storeName: {
    marginBottom: 10,
    fontWeight: 'bold',
    fontSize: 20,
  },
  waiting: {
    marginBottom: 10,
    fontWeight: 'bold',
    fontSize: 16,
  },
  waitingFor: {
    fontWeight: 'bold',
    fontSize: 20,
  },
  actionGroup: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginTop: 20,
  },
  invisible: {
    opacity: 0,
    height: 0,
    width: 80,
  },
  cancelSuggest: {
    textAlign: 'right',
    width: 80,
    color: '#666',
    fontSize: 14,
  },
  row: {
    flexDirection: 'row',
    marginBottom: 15,
  },
  name: {
    flex: 2,
    fontWeight: 'bold',
    fontSize: 16,
  },
  hour: {
    flex: 1,
    fontWeight: 'bold',
    fontSize: 16,
  },
  distance: {
    flex: 1,
    fontWeight: 'bold',
    fontSize: 16,
  },
  collapseWrapper: {
    flexDirection: 'row',
    justifyContent: 'flex-end',
  },
  collapse: {
    color: '#666',
  },
});

export default StoreInfo;
