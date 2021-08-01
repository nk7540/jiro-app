import React, {useCallback} from 'react';
import {launchCamera} from 'react-native-image-picker';
import {useNavigation, useFocusEffect} from '@react-navigation/native';
import {useAppDispatch, setPost, useAppSelector, increment, updateConsumingFor} from 'services/store';
import {SafeAreaView, View, StyleSheet} from 'react-native';
import {Text, Button} from 'react-native-elements';
import {secToHour} from 'utils';
import ImageList from 'components/ImageList';
import {ScreenNames} from 'services/navigation';

const CreateImages = () => {
  const navigation = useNavigation();
  const {images, consumingFor} = useAppSelector(state => state.post);
  const dispatch = useAppDispatch();

  useFocusEffect(
    useCallback(() => {
      let intervalId: NodeJS.Timer;
      launchCamera({mediaType: 'photo'}, res => {
        if (res.errorCode) {
          console.log(res.errorCode, res.errorMessage);
        } else if (res.assets) {
          dispatch(
            setPost({
              key: 'images',
              value: [
                {
                  uri: res.assets[0].uri || '',
                  name: `${Date.now()}.jpg`,
                  type: res.assets[0].type || 'image/jpeg',
                },
              ],
            }),
          );
        }

        dispatch(setPost({key: 'consumingSince', value: new Date().getTime()}));
        intervalId = setInterval(() => {
          dispatch(updateConsumingFor(new Date().getTime()));
        }, 1000);
      });

      return () => clearInterval(intervalId);
    }, [dispatch]),
  );

  return (
    <SafeAreaView>
      <ImageList images={images} />
      <View style={styles.container}>
        <View style={styles.consuming}>
          <Text style={styles.consumingFor}>{secToHour(consumingFor)}</Text>
        </View>
        <Button title="完食" onPress={() => navigation.navigate(ScreenNames.CreatePost)} />
      </View>
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  container: {
    marginHorizontal: 25,
  },
  consuming: {
    alignItems: 'center',
    padding: 20,
    borderRadius: 10,
    marginTop: 10,
    marginBottom: 20,
    backgroundColor: '#E0E5E3',
  },
  consumingFor: {
    fontWeight: 'bold',
    fontSize: 60,
  },
});

export default CreateImages;
