import React, {ReactNode, useState, useCallback} from 'react';
import {Text, View, ScrollView, StyleSheet, RefreshControl, FlatList} from 'react-native';

const styles = StyleSheet.create({
  serverError: {
    justifyContent: 'center',
    height: '100%',
  },
  serverErrorMsg: {
    color: 'grey',
  },
});

interface Props<T> {
  loading: boolean;
  error: string | undefined;
  objects: T[] | undefined;
  horizontal?: boolean;
  refetch?: (args?: any) => Promise<any>;
  children: (object: T) => ReactNode;
}

interface IObject {
  id: number;
}

const ListWithError = <T extends IObject>({loading, error, objects, horizontal, refetch, children}: Props<T>) => {
  const [refreshing, setRefreshing] = useState(false);

  const onRefresh = useCallback(() => {
    setRefreshing(true);
    refetch && refetch().then(() => setRefreshing(false));
  }, [refetch]);

  if (loading) {
    return <></>;
  } else if (error || !objects) {
    return (
      <View style={styles.serverError}>
        <Text style={styles.serverErrorMsg}>サーバーエラーです。時間を置いてから再度お試しください。</Text>
      </View>
    );
  } else if (objects.length === 0) {
    return (
      <View style={styles.serverError}>
        <Text style={styles.serverErrorMsg}>データがありません。</Text>
      </View>
    );
  } else {
    return (
      <FlatList
        data={objects}
        renderItem={({item}) => <>{children(item)}</>}
        horizontal={horizontal}
        refreshControl={refetch && <RefreshControl refreshing={refreshing} onRefresh={onRefresh} />}
      />
    );
  }
};

export default ListWithError;
