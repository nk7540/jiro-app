import React from 'react';
import {ScrollView, Image} from 'react-native';
import {ReactNativeFile} from 'apollo-upload-client';

interface Props {
  images: ReactNativeFile[];
}

const ImageList = ({images}: Props) => {
  return (
    <ScrollView horizontal>
      {images.map(image => (
        <Image key={image.name} source={{uri: image.uri}} />
      ))}
    </ScrollView>
  );
};

export default ImageList;
