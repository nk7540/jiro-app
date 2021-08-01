/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL fragment: post
// ====================================================

export interface post_user {
  __typename: "User";
  id: number;
  nickname: string;
}

export interface post_images {
  __typename: "PostImage";
  id: number;
  filename: string;
}

export interface post_tags {
  __typename: "Tag";
  id: number;
  name: string;
  kind: string;
  storeId: number | null;
}

export interface post {
  __typename: "Post";
  id: number;
  user: post_user;
  comment: string;
  images: post_images[];
  tags: post_tags[];
  waitingFor: number | null;
  consumingFor: number | null;
}
