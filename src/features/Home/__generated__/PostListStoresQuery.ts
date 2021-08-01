/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: PostListStoresQuery
// ====================================================

export interface PostListStoresQuery_stores {
  __typename: "Store";
  id: number;
  name: string;
  latitude: number;
  longitude: number;
}

export interface PostListStoresQuery {
  stores: PostListStoresQuery_stores[];
}
