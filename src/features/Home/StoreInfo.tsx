import React, {useState, useCallback} from 'react';
import {View, StyleSheet} from 'react-native';
import {Text, Button} from 'react-native-elements';
import {secToHour} from 'utils';
import {useAppDispatch, useAppSelector, setPost, updateWaitingFor, resetWaiting} from 'services/store';
import {useNavigation} from '@react-navigation/native';
import {ScreenNames} from 'services/navigation';

interface Props {
  setNoUpdate: (noUpdate: boolean) => void;
}

const StoreInfo = ({setNoUpdate}: Props) => {
  const [waitingIntervalId, setWaitingIntervalId] = useState<NodeJS.Timer>();
  const navigation = useNavigation();
  const post = useAppSelector(state => state.post);
  const dispatch = useAppDispatch();

  const lineUp = () => {
    dispatch(setPost({key: 'waitingSince', value: new Date().getTime()}));
    const intervalId = setInterval(() => {
      dispatch(updateWaitingFor(new Date().getTime()));
    }, 1000);
    setWaitingIntervalId(intervalId);
  };

  const cancelSuggest = () => {
    setNoUpdate(true);
    setTimeout(() => setNoUpdate(false), 60000);
    dispatch(setPost({key: 'store', value: {id: 0, name: ''}}));
  };

  const clearWaiting = useCallback(() => {
    dispatch(resetWaiting());
    clearInterval(waitingIntervalId as NodeJS.Timer);
  }, [dispatch, waitingIntervalId]);

  const orderArrived = () => {
    clearWaiting();
    navigation.navigate(ScreenNames.CreateImages);
  };

  if (post.waitingFor > 0) {
    return (
      <View style={styles.store}>
        <Text style={styles.storeName}>{post.store.name}</Text>
        <Text style={styles.waiting}>
          待機中...<Text style={styles.waitingFor}>{secToHour(post.waitingFor)}</Text>
        </Text>
        <View style={styles.actionGroup}>
          <View style={styles.invisible} />
          <Button title="着丼" onPress={orderArrived} />
          <Text style={styles.cancelSuggest} onPress={clearWaiting}>
            キャンセル
          </Text>
        </View>
      </View>
    );
  } else if (post.store.id > 0) {
    return (
      <View style={styles.store}>
        <Text>
          今<Text style={styles.storeName}>{post.store.name}</Text>の近くにいますか？
        </Text>
        <View style={styles.actionGroup}>
          <View style={styles.invisible} />
          <Button title="並ぶ" onPress={lineUp} />
          <Text style={styles.cancelSuggest} onPress={cancelSuggest}>
            いいえ
          </Text>
        </View>
      </View>
    );
  } else {
    return null;
  }
};

const styles = StyleSheet.create({
  store: {
    padding: 20,
    borderRadius: 20,
    marginHorizontal: 25,
    marginTop: 10,
    marginBottom: 20,
    backgroundColor: '#E0E5E3',
  },
  storeName: {
    marginBottom: 10,
    fontWeight: 'bold',
    fontSize: 20,
  },
  waiting: {
    marginBottom: 10,
    fontWeight: 'bold',
    fontSize: 16,
  },
  waitingFor: {
    fontWeight: 'bold',
    fontSize: 20,
  },
  actionGroup: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginTop: 20,
  },
  invisible: {
    opacity: 0,
    height: 0,
    width: 80,
  },
  cancelSuggest: {
    textAlign: 'right',
    width: 80,
    color: '#666',
    fontSize: 14,
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
  collapseWrapper: {
    flexDirection: 'row',
    justifyContent: 'flex-end',
  },
  collapse: {
    color: '#666',
  },
});

export default StoreInfo;
