import React from 'react';
import {gql, useQuery, useMutation} from '@apollo/client';
import Tag from 'components/Tag';
import {useAppSelector, useAppDispatch, setPost} from 'services/store';
import {SafeAreaView} from 'react-native';
import ImageList from 'components/ImageList';
import {Text, Input, Button} from 'react-native-elements';
import {secToHour} from 'utils';
import {useNavigation, StackActions} from '@react-navigation/native';
import {TagsQuery} from './__generated__/TagsQuery';
import {CreatePostMutation} from './__generated__/CreatePostMutation';

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

const CreatePost = () => {
  const navigation = useNavigation();
  const {} = useQuery<TagsQuery>(TAGS_QUERY);
  const [createPost] = useMutation<CreatePostMutation>(CREATE_POST_MUTATION, {
    onCompleted: () => navigation.dispatch(StackActions.popToTop),
  });
  const post = useAppSelector(state => state.post);
  const dispatch = useAppDispatch();

  return (
    <SafeAreaView>
      <ImageList images={post.images} />
      <Text>待ち時間：{secToHour(post.waitingFor)}</Text>
      <Text>完食時間：{secToHour(post.consumingFor)}</Text>
      <Input value={post.comment} onChangeText={v => dispatch(setPost({key: 'comment', value: v}))} />
      <Button title="投稿する" onPress={() => createPost({variables: post})} />
    </SafeAreaView>
  );
};

export default CreatePost;
