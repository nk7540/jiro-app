/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: PostsQuery
// ====================================================

export interface PostsQuery_posts_user {
  __typename: "User";
  id: number;
  nickname: string;
}

export interface PostsQuery_posts_images {
  __typename: "PostImage";
  id: number;
  filename: string;
}

export interface PostsQuery_posts_tags {
  __typename: "Tag";
  id: number;
  name: string;
  kind: string;
  storeId: number | null;
}

export interface PostsQuery_posts {
  __typename: "Post";
  id: number;
  user: PostsQuery_posts_user;
  comment: string;
  images: PostsQuery_posts_images[];
  tags: PostsQuery_posts_tags[];
  waitingFor: number | null;
  consumingFor: number | null;
}

export interface PostsQuery {
  posts: PostsQuery_posts[];
}
