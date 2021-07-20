import React, {useEffect} from 'react';
import {launchImageLibrary} from 'react-native-image-picker';
import {useNavigation} from '@react-navigation/native';
import {useAppDispatch, setPost, useAppSelector, increment} from 'services/store';
import {ReactNativeFile} from 'apollo-upload-client';
import {SafeAreaView, ScrollView, Image, View} from 'react-native';
import {Text} from 'react-native-elements';
import {secToHour} from 'utils';

const CreateImages = () => {
  const navigation = useNavigation();
  const {images, consumingFor} = useAppSelector(state => state.post);
  const dispatch = useAppDispatch();

  useEffect(() => {
    let intervalId: NodeJS.Timer;
    launchImageLibrary({mediaType: 'photo', selectionLimit: 99}, res => {
      if (res.errorCode) {
        console.log(res.errorCode, res.errorMessage);
        return;
      } else if (res.didCancel) {
        navigation.goBack();
        return;
      } else if (res.assets) {
        dispatch(
          setPost({
            key: 'images',
            value: res.assets.map(asset => {
              return new ReactNativeFile({
                uri: asset.uri || '',
                name: `${Date.now()}.jpg`,
                type: asset.type,
              });
            }),
          }),
        );

        intervalId = setInterval(() => {
          dispatch(increment('consumingFor'));
        }, 1000);
      }
    });

    return () => clearInterval(intervalId);
  });

  return (
    <SafeAreaView>
      <ScrollView horizontal>
        {images.map(image => (
          <Image source={{uri: image.uri}} />
        ))}
      </ScrollView>
      <View>
        <Text>{secToHour(consumingFor)}</Text>
      </View>
    </SafeAreaView>
  );
};

export default CreateImages;
