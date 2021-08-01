/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

//==============================================================
// START Enums and Input Objects
//==============================================================

export enum HourStatus {
  Closed = "Closed",
  ClosingSoon = "ClosingSoon",
  Open = "Open",
  OpeningSoon = "OpeningSoon",
}

export interface CreatePost {
  images: any[];
  comment: string;
  waitingFor: number;
  consumingFor: number;
  tagsIds: number[];
  status: string;
}

export interface CreateUser {
  nickname: string;
  email: string;
  password: string;
  passwordConfirmation: string;
}

//==============================================================
// END Enums and Input Objects
//==============================================================
