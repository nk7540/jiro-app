import React, {useState} from 'react';
import {Image, StyleSheet, Dimensions, View, ListRenderItem} from 'react-native';
import {ReactNativeFile} from 'apollo-upload-client';
import Carousel, {Pagination} from 'react-native-snap-carousel';
import {Text} from 'react-native-elements';

const {width: screenWidth} = Dimensions.get('window');

interface Props {
  images: ReactNativeFile[];
}

const ImageList = ({images}: Props) => {
  const [activeSlide, setActiveSlide] = useState(0);

  const renderItem: ListRenderItem<ReactNativeFile> = ({item}) => {
    return (
      <View style={styles.item}>
        <Image source={{uri: item.uri}} style={styles.image} />
      </View>
    );
  };

  if (images.length === 0) {
    return (
      <View style={styles.noImage}>
        <Text style={styles.noImageText}>写真がありません。</Text>
      </View>
    );
  }

  return (
    <Carousel
      sliderWidth={screenWidth}
      sliderHeight={screenWidth}
      itemWidth={screenWidth - 60}
      data={images}
      renderItem={renderItem}
      onSnapToItem={i => setActiveSlide(i)}>
      <Pagination
        dotsLength={images.length}
        activeDotIndex={activeSlide}
        dotStyle={styles.dot}
        inactiveDotOpacity={0.4}
        inactiveDotScale={0.6}
      />
    </Carousel>
  );
};

const styles = StyleSheet.create({
  noImage: {
    justifyContent: 'center',
    alignItems: 'center',
    width: screenWidth,
    height: screenWidth,
    backgroundColor: '#EAEFED',
  },
  noImageText: {
    color: '#666',
  },
  item: {
    width: screenWidth,
    height: screenWidth,
  },
  image: {
    ...StyleSheet.absoluteFillObject,
    resizeMode: 'cover',
    borderRadius: 8,
  },
  dot: {
    width: 10,
    height: 10,
    borderRadius: 5,
    marginHorizontal: 8,
    backgroundColor: 'rgba(255, 255, 255, 0.92)',
  },
});

export default ImageList;
