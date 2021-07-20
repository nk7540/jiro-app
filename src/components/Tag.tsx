import React from 'react';
import {FragC} from 'services/graphql';
import {View, StyleSheet, TouchableOpacity} from 'react-native';
import {Text} from 'react-native-elements';
import {gql} from '@apollo/client';

interface FragmentProps {
  data: any;
}

interface Props {
  onPress: () => void;
}

const Tag: FragC<FragmentProps, Props> = ({data, onPress}) => {
  return (
    <TouchableOpacity onPress={onPress}>
      <View style={styles.wrapper}>
        <Text style={styles.name}>{data.name}</Text>
      </View>
    </TouchableOpacity>
  );
};

Tag.fragments = {
  data: gql`
    fragment tag on Tag {
      id
      name
      kind
      storeId
    }
  `,
};

const styles = StyleSheet.create({
  wrapper: {
    justifyContent: 'center',
    alignItems: 'center',
    padding: 10,
    borderRadius: 100,
    backgroundColor: '#BBB',
  },
  name: {
    fontWeight: 'bold',
  },
});

export default Tag;
