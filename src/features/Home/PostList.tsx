import React from 'react';
import {gql, useQuery} from '@apollo/client';
import PostListItem from './PostListItem';
import {PostsQuery} from './__generated__/PostsQuery';
import ListWithError from 'components/ListWithError';

const POSTS_QUERY = gql`
  query PostsQuery {
    posts {
      ...post
    }
  }
  ${PostListItem.fragments.data}
`;

const PostList = () => {
  const {loading, error, data, refetch} = useQuery<PostsQuery>(POSTS_QUERY);

  return (
    <ListWithError loading={loading} error={error?.message} objects={data?.posts} refetch={refetch}>
      {post => <PostListItem data={post} />}
    </ListWithError>
  );
};

export default PostList;
