import React from 'react';
import {FragC} from 'services/graphql';
import {gql} from '@apollo/client';
import {store} from './__generated__/store';
import {View, StyleSheet} from 'react-native';
import {Text} from 'react-native-elements';
import {secToHour} from 'utils';

interface FragmentProps {
  data: store;
}

interface Props {
  distance: number;
  setCurrentStore: () => void;
}

const StoreItem: FragC<FragmentProps, Props> = ({data, distance, setCurrentStore}) => {
  return (
    <View style={styles.row}>
      <Text style={styles.name} onPress={setCurrentStore}>
        {data.name}
      </Text>
      <Text style={styles.hour}>
        {secToHour(data.closeIn, {sec: false}) + (data.closeIn === 0 && data.note ? `(${data.note})` : '')}
      </Text>
      <Text style={styles.distance}>{distance ? Math.round(distance) : '   -   '} km</Text>
    </View>
  );
};

StoreItem.fragments = {
  data: gql`
    fragment store on Store {
      id
      name
      closeIn
      note
    }
  `,
};

const styles = StyleSheet.create({
  row: {
    flexDirection: 'row',
    marginBottom: 10,
    height: 25,
  },
  name: {
    flex: 2,
    fontWeight: 'bold',
    fontSize: 20,
  },
  hour: {
    flex: 1,
    fontSize: 16,
  },
  distance: {
    flex: 1,
    fontSize: 16,
  },
});

export default StoreItem;
