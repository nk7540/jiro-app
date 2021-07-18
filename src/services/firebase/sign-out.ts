import auth from '@react-native-firebase/auth';

export default async function signOut(): Promise<void> {
  await auth().signOut();
}
