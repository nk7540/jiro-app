import React, {useState, useCallback, useMemo} from 'react';
import {gql, useQuery, useMutation} from '@apollo/client';
import Tag from 'components/Tag';
import {useAppSelector, useAppDispatch, setPost, addTagId, removeTagId, resetPost} from 'services/store';
import {SafeAreaView, StyleSheet, View, ScrollView} from 'react-native';
import ImageList from 'components/ImageList';
import {Text, Input, Button, Switch} from 'react-native-elements';
import {secToHour} from 'utils';
import {useNavigation, StackActions} from '@react-navigation/native';
import {TagsQuery, TagsQuery_tags} from './__generated__/TagsQuery';
import {CreatePostMutation} from './__generated__/CreatePostMutation';
import {CreatePost} from '../../../__generated__/globalTypes';

const TAGS_QUERY = gql`
  query TagsQuery {
    tags {
      ...tag
    }
  }
  ${Tag.fragments.data}
`;

const CREATE_POST_MUTATION = gql`
  mutation CreatePostMutation($input: CreatePost!) {
    createPost(input: $input)
  }
`;

const CreatePostIndex = () => {
  const [storeVisible, setStoreVisible] = useState(false);
  const navigation = useNavigation();
  const {data} = useQuery<TagsQuery>(TAGS_QUERY);
  const [createPost] = useMutation<CreatePostMutation>(CREATE_POST_MUTATION, {
    onCompleted: () => {
      dispatch(resetPost());
      navigation.dispatch(StackActions.popToTop);
    },
  });
  const post = useAppSelector(state => state.post);
  const dispatch = useAppDispatch();

  const tags = useMemo(() => {
    let currentStoreTag = {} as TagsQuery_tags;
    let storeTags = [] as TagsQuery_tags[];
    let toppingTags = [] as TagsQuery_tags[];
    if (data) {
      storeTags = data.tags.filter(tag => tag.kind === 'store');
      const currentStoreTagIndex = storeTags.findIndex(tag => tag.id === post.store.id);
      currentStoreTag = storeTags.splice(currentStoreTagIndex, 1)[0];
      toppingTags = data.tags.filter(tag => tag.kind === 'topping');
    }

    return {currentStoreTag, storeTags, toppingTags};
  }, [data, post.store.id]);

  const isActive = useCallback(
    (id: number) => {
      return post.tagIds.includes(id);
    },
    [post.tagIds],
  );

  const onPressTag = useCallback(
    (id: number) => {
      dispatch(isActive(id) ? removeTagId(id) : addTagId(id));
    },
    [isActive, dispatch],
  );

  const _createPost = useCallback(() => {
    const tagIds = [...post.tagIds, tags.currentStoreTag.id];
    const input: CreatePost = {
      images: post.images,
      comment: post.comment,
      waitingFor: post.waitingFor,
      consumingFor: post.consumingFor,
      tagsIds: tagIds,
      status: post.status,
    };
    createPost({variables: {input: input}});
  }, [post, tags.currentStoreTag, createPost]);

  return (
    <SafeAreaView>
      <ScrollView>
        <ImageList images={post.images} />
        <Text style={styles.periodWrapper}>
          待ち時間：<Text style={styles.period}>{secToHour(post.waitingFor)}</Text>
        </Text>
        <Text style={styles.periodWrapper}>
          完食時間：<Text style={styles.period}>{secToHour(post.consumingFor)}</Text>
        </Text>
        <View style={styles.container}>
          <View style={styles.tagLists}>
            <View style={styles.storeTagList}>
              <View style={styles.currentStoreTag}>
                <Tag
                  data={tags.currentStoreTag}
                  onPress={() =>
                    dispatch(
                      setPost({
                        key: 'store',
                        value: {id: tags.currentStoreTag.storeId as number, name: tags.currentStoreTag.name},
                      }),
                    )
                  }
                  active={tags.currentStoreTag.storeId === post.store.id}
                />
                <Text style={styles.action} onPress={() => setStoreVisible(true)}>
                  他の店舗を表示する
                </Text>
              </View>
              <View style={styles.tagList}>
                {storeVisible &&
                  tags.storeTags.map(tag => (
                    <Tag
                      data={tag}
                      onPress={() =>
                        dispatch(setPost({key: 'store', value: {id: tag.storeId as number, name: tag.name}}))
                      }
                      active={tag.storeId === post.store.id}
                    />
                  ))}
              </View>
            </View>
            <View style={styles.tagList}>
              {tags.toppingTags.map(tag => (
                <Tag data={tag} onPress={() => onPressTag(tag.id)} active={isActive(tag.id)} />
              ))}
            </View>
          </View>
          <Input
            style={styles.input}
            inputContainerStyle={styles.inputContainer}
            placeholder="コメントを入力してください。"
            multiline={true}
            numberOfLines={5}
            value={post.comment}
            onChangeText={v => dispatch(setPost({key: 'comment', value: v}))}
          />
          <View style={styles.status}>
            <Text>公開する</Text>
            <Switch
              value={post.status === 'public'}
              onChange={() => {
                dispatch(setPost({key: 'status', value: post.status === 'public' ? 'private' : 'public'}));
              }}
            />
          </View>
          <Button title="投稿する" onPress={_createPost} />
        </View>
      </ScrollView>
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  periodWrapper: {
    alignSelf: 'center',
    marginBottom: 10,
  },
  period: {
    fontWeight: 'bold',
    fontSize: 20,
  },
  container: {
    marginHorizontal: 25,
  },
  tagLists: {
    padding: 20,
    borderRadius: 20,
    marginTop: 10,
    backgroundColor: '#EAEFED',
  },
  storeTagList: {
    borderBottomWidth: 1,
    borderBottomColor: '#666',
    marginBottom: 10,
  },
  currentStoreTag: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: 10,
  },
  action: {
    marginBottom: 10,
    color: '#666',
  },
  tagList: {
    flexDirection: 'row',
    flexWrap: 'wrap',
  },
  input: {
    fontSize: 14,
  },
  inputContainer: {
    height: 120,
  },
  status: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: 20,
  },
});

export default CreatePostIndex;
