/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

import { HourStatus } from "./../../../../__generated__/globalTypes";

// ====================================================
// GraphQL fragment: store
// ====================================================

export interface store {
  __typename: "Store";
  id: number;
  name: string;
  hourStatus: HourStatus;
  nextHour: string;
  note: string;
}
