/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: TagsQuery
// ====================================================

export interface TagsQuery_tags {
  __typename: "Tag";
  id: number;
  name: string;
  kind: string;
  storeId: number | null;
}

export interface TagsQuery {
  tags: TagsQuery_tags[];
}
