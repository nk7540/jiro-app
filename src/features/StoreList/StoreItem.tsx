import React from 'react';
import {FragC} from 'services/graphql';
import {gql} from '@apollo/client';
import {View, StyleSheet} from 'react-native';
import {Text} from 'react-native-elements';
import {secToHour} from 'utils';
import {store} from './__generated__/store';
import {HourStatus} from '../../../__generated__/globalTypes';

interface FragmentProps {
  data: store;
}

interface Props {
  distance: number;
}

const StoreItem: FragC<FragmentProps, Props> = ({data, distance}) => {
  const hourText = (hourStatus: HourStatus) => {
    switch (hourStatus) {
      case HourStatus.Closed:
        return (
          <Text>
            <Text style={styles.closed}>閉店中 </Text>開店時間：
          </Text>
        );
      case HourStatus.OpeningSoon:
        return <Text style={styles.soon}>まもなく営業開始：</Text>;
      case HourStatus.Open:
        return (
          <Text>
            <Text style={styles.open}>営業中 </Text>営業終了：
          </Text>
        );
      case HourStatus.ClosingSoon:
        return <Text style={styles.soon}>まもなく営業終了：</Text>;
    }
  };

  return (
    <View style={styles.row}>
      <Text style={styles.name}>{data.name}</Text>
      <Text style={styles.hour}>
        {hourText(data.hourStatus)}
        <Text>{data.nextHour.substring(-1, 5)}</Text>
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
      hourStatus
      nextHour
      note
    }
  `,
};

const styles = StyleSheet.create({
  row: {
    flexDirection: 'row',
    alignItems: 'center',
    padding: 20,
    borderRadius: 20,
    marginHorizontal: 25,
    marginBottom: 20,
    backgroundColor: '#fff',
    shadowColor: '#000',
    shadowOffset: {
      width: 0,
      height: 2,
    },
    shadowOpacity: 0.25,
    shadowRadius: 3.84,
  },
  name: {
    flex: 3,
    fontWeight: 'bold',
    fontSize: 20,
    color: '#000',
  },
  hour: {
    flex: 3,
    fontSize: 14,
  },
  open: {
    color: '#3c7e40',
  },
  closed: {
    color: '#c84031',
  },
  soon: {
    color: '#debd96',
  },
  distance: {
    flex: 1,
    fontSize: 16,
  },
});

export default StoreItem;
