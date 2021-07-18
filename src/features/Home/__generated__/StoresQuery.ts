/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: StoresQuery
// ====================================================

export interface StoresQuery_stores {
  __typename: "Store";
  id: number;
  name: string;
  closeAfter: number;
  latitude: number;
  longitude: number;
}

export interface StoresQuery {
  stores: StoresQuery_stores[];
}
