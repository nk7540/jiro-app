import auth, {FirebaseAuthTypes} from '@react-native-firebase/auth';
import {AuthUser, AuthUserError} from './authUser';

export default async function signInWithPassword(email: string, password: string): Promise<AuthUser> {
  return await auth()
    .signInWithEmailAndPassword(email, password)
    .then(async (res: FirebaseAuthTypes.UserCredential) => {
      if (!res.user) {
        throw new AuthUserError('Eメールが登録されていないか、パスワードが間違っています。もう一度お試しください。');
      }

      const authUser: AuthUser = {
        uid: res.user.uid,
        email: res.user.email || '',
        emailVerified: res.user.emailVerified,
        token: '',
        creationTime: res.user.metadata?.creationTime,
        lastSignInTime: res.user.metadata?.lastSignInTime,
      };

      await auth()
        .currentUser?.getIdToken()
        .then((token: string) => {
          authUser.token = token;
        });

      return authUser;
    })
    .catch(err => {
      console.log(err);
      throw new AuthUserError('Eメールが登録されていないか、パスワードが間違っています。もう一度お試しください。');
    });
}
