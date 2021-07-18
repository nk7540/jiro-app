import React from 'react';
import {FragC} from 'services/graphql';
import {gql} from '@apollo/client';
import {store} from './__generated__/store';
import {View} from 'react-native';
import {Text} from 'react-native-elements';

interface FragmentProps {
  data: store;
}

interface Props {
  distance: number;
}

const StoreItem: FragC<FragmentProps, Props> = ({data, distance}) => {
  return (
    <View>
      <Text>{data.name}</Text>
      <Text>{data.closeAfter}</Text>
      <Text>{distance}</Text>
    </View>
  );
};

StoreItem.fragments = {
  data: gql`
    fragment store on Store {
      id
      name
      closeAfter
    }
  `,
};

export default StoreItem;
