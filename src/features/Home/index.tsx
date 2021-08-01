import React, {FC, useState} from 'react';
import {SafeAreaView, ScrollView, StyleSheet} from 'react-native';
import StoreInfo from './StoreInfo';
import PostList from './PostList';

const Home: FC = () => {
  const [noUpdate, setNoUpdate] = useState(false);

  return (
    <SafeAreaView style={styles.flex1}>
      <StoreInfo setNoUpdate={setNoUpdate} />
      <PostList noUpdate={noUpdate} />
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  flex1: {
    flex: 1,
  },
  storeInfo: {
    paddingBottom: 65,
  },
});

export default Home;
