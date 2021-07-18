export interface AuthUser {
  uid: string;
  email: string;
  emailVerified: boolean;
  token: string;
  creationTime?: string;
  lastSignInTime?: string;
}

export class AuthUserError extends Error {}
