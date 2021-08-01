import React, {useEffect, useState} from 'react';
import {FragC} from 'services/graphql';
import {gql} from '@apollo/client';
import Tag from 'components/Tag';
import {post} from './__generated__/post';
import {View, StyleSheet} from 'react-native';
import storage from '@react-native-firebase/storage';
import ImageList from 'components/ImageList';
import {Text} from 'react-native-elements';
import {secToHour} from 'utils';

interface FragmentProps {
  data: post;
}

const PostListItem: FragC<FragmentProps> = ({data}) => {
  // @TODO implementation
  // const onPressTag = (tag: post_tags) => {
  // };
  const [images, setImages] = useState([] as {id: number; uri: string}[]);

  useEffect(() => {
    const _getImages = Promise.all(
      data.images.map(async image => {
        const uri = await storage().ref(image.filename).getDownloadURL();
        return {id: image.id, uri};
      }),
    );
    _getImages.then(_images => setImages(_images));
  }, [data.images, setImages]);

  return (
    <View>
      <ImageList images={images} />
      <View style={styles.postData}>
        <View style={styles.tagList}>
          {data.tags.map(tag => (
            <Tag key={tag.id} data={tag} onPress={() => tag} />
          ))}
        </View>
        <Text style={styles.comment}>
          <Text style={styles.nickname}>{data.user.nickname} </Text>
          {data.comment}
        </Text>
        <View style={styles.periodWrapper}>
          <Text style={styles.period}>待ち時間：{secToHour(data.waitingFor as number)}</Text>
          <Text style={styles.period}>完食：{secToHour(data.consumingFor as number)}</Text>
        </View>
      </View>
    </View>
  );
};

PostListItem.fragments = {
  data: gql`
    fragment post on Post {
      id
      user {
        id
        nickname
      }
      comment
      images {
        id
        filename
      }
      tags {
        ...tag
      }
      waitingFor
      consumingFor
    }
    ${Tag.fragments.data}
  `,
};

const styles = StyleSheet.create({
  postData: {
    marginHorizontal: 10,
    marginBottom: 35,
  },
  tagList: {
    flexDirection: 'row',
    flexWrap: 'wrap',
    marginVertical: 10,
  },
  comment: {
    marginBottom: 10,
  },
  nickname: {
    paddingRight: 10,
    fontWeight: 'bold',
  },
  periodWrapper: {
    flexDirection: 'row',
  },
  period: {
    marginRight: 10,
    color: '#666',
  },
});

export default PostListItem;
