import React, {useCallback} from 'react';
import {SafeAreaView} from 'react-native';
import ListWithError from 'components/ListWithError';
import StoreItem from './StoreItem';
import {useFocusEffect} from '@react-navigation/native';
import {useQuery, gql} from '@apollo/client';
import {useAppSelector} from 'services/store';
import {StoresQuery} from './__generated__/StoresQuery';

const STORES_QUERY = gql`
  query StoresQuery {
    stores {
      ...store
    }
  }
  ${StoreItem.fragments.data}
`;

const StoreList = () => {
  const {loading, error, data, refetch} = useQuery<StoresQuery>(STORES_QUERY);
  const distances = useAppSelector(state => state.distances);

  useFocusEffect(
    useCallback(() => {
      refetch();
      const intervalId = setInterval(() => {
        refetch();
      }, 60000);
      return () => clearInterval(intervalId);
      // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []),
  );

  return (
    <SafeAreaView>
      <ListWithError loading={loading} error={error?.message} objects={data?.stores}>
        {store => <StoreItem data={store} distance={distances[store.id]} />}
      </ListWithError>
    </SafeAreaView>
  );
};

export default StoreList;
