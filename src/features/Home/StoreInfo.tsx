import React, {useEffect, useState, useCallback} from 'react';
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
import {useNavigation} from '@react-navigation/native';
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

  useEffect(() => {
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
        if (key === currentUserId) {
          dispatch(setPost({key: 'store', value: store}));
        }
      });
      query.on('key_exited', (key: string) => {
        console.log('someone exited');
        if (key === currentUserId) {
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
  }, [currentUserData, data, updateDistances, dispatch, waitingIntervalId]);

  useEffect(() => {
    const intervalId = setInterval(() => {
      refetch();
    }, 60000);
    return () => clearInterval(intervalId);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

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

  const orderArrived = () => {
    navigation.navigate(ScreenNames.CreateImages);
  };

  if (post.waitingFor > 0) {
    return (
      <>
        <Text>{post.store.name}</Text>
        <Text>待機中</Text>
        <Text>{secToHour(post.waitingFor)}</Text>
        <Button title="着丼" onPress={orderArrived} />
        <Button title="キャンセル" onPress={getOut} />
      </>
    );
  } else if (post.store.id > 0) {
    return (
      <View style={styles.store}>
        <Text>今{post.store.name}の近くにいますか？</Text>
        <View style={styles.actionGroup}>
          <View style={styles.invisible} />
          <Button style={styles.lineUp} title="並ぶ" onPress={lineUp} />
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
    );
  }
};

const styles = StyleSheet.create({
  store: {
    padding: 20,
    borderRadius: 20,
    marginHorizontal: 25,
    marginTop: 10,
    backgroundColor: '#E9BF40',
  },
  actionGroup: {
    flexDirection: 'row',
    justifyContent: 'space-between',
  },
  invisible: {
    opacity: 0,
    height: 0,
  },
  lineUp: {
    paddingHorizontal: 30,
    paddingVertical: 15,
    borderRadius: 20,
    marginBottom: 30,
  },
  cancelSuggest: {
    color: '#999',
    fontSize: 16,
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
});

export default StoreInfo;
