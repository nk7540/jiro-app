/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

import { HourStatus } from "./../../../../__generated__/globalTypes";

// ====================================================
// GraphQL query operation: StoresQuery
// ====================================================

export interface StoresQuery_stores {
  __typename: "Store";
  id: number;
  name: string;
  hourStatus: HourStatus;
  nextHour: string;
  note: string;
}

export interface StoresQuery {
  stores: StoresQuery_stores[];
}
