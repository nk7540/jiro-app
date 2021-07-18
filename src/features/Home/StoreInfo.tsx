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
  const [currentStore, setCurrentStore] = useState({id: 0, name: ''});
  const [waitingFor, setWaitingFor] = useState(0);
  const {loading, error, data, refetch} = useQuery<StoresQuery>(STORES_QUERY);
  const {data: currentUserData} = useQuery<StoreInfoUserQuery>(STORE_INFO_USER_QUERY);

  const updateDistances = useCallback(
    (coords: {latitude: number; longitude: number}, stores: StoresQuery_stores[]) => {
      let newDistances = Object.assign({}, distances);
      stores.forEach(store => {
        const distance = distanceBetween([coords.latitude, coords.longitude], [store.latitude, store.longitude]);
        newDistances = {...newDistances, [store.id]: distance};
      });
      setDistances(newDistances);
    },
    [distances],
  );

  useEffect(() => {
    if (!currentUserData || !data) {
      return;
    }
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
      query.on('key_entered', () => {
        console.log('someone entered');
        setCurrentStore(c => ({...c, id: store.id}));
      });
      query.on('key_exited', () => {
        console.log('someone exited');
        setCurrentStore({id: 0, name: ''});
      });
      cancels.push(query.cancel.bind(query));
    });

    return () => {
      Geolocation.clearWatch(watchId);
      cancels.forEach(cancel => cancel());
    };
  }, [currentUserData, data, updateDistances]);

  useEffect(() => {
    const intervalId = setInterval(() => {
      refetch();
    }, 60000);
    return () => clearInterval(intervalId);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const lineUp = () => {
    setInterval(() => {
      setWaitingFor(waitingFor => waitingFor + 1);
    }, 1000);
  };

  const orderArrived = () => {
    // @TODO navigate to create post page
  };

  console.log(distances);

  if (waitingFor > 0) {
    return (
      <>
        <Text>{currentStore.name}</Text>
        <Text>待機中</Text>
        <Text>{new Date(waitingFor * 1000).toISOString().substr(11, 8)}</Text>
        <Button title="着丼" onPress={orderArrived} />
      </>
    );
  } else if (currentStore.id > 0) {
    return (
      <View>
        <Text>今{currentStore.name}の近くにいますか？</Text>
        <Button title="並ぶ" onPress={lineUp} />
      </View>
    );
  } else {
    return (
      <View style={styles.store}>
        <ListWithError loading={loading} error={error?.message} objects={data?.stores}>
          {store => <StoreItem data={store} distance={distances[store.id]} />}
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
  },
});

export default StoreInfo;
