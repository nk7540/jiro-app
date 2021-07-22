import React, {FC} from 'react';
import {SafeAreaView, ScrollView} from 'react-native';
import StoreInfo from './StoreInfo';
import PostList from './PostList';

const Home: FC = () => {
  return (
    <SafeAreaView>
      <ScrollView>
        <StoreInfo />
      </ScrollView>
      <PostList />
    </SafeAreaView>
  );
};

export default Home;
