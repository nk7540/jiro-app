import {StackScreenProps} from '@react-navigation/stack';

export type StackParamList = {
  Loading: undefined;
  SignUp: undefined;
  SignIn: undefined;
  Home: undefined;
};

export type ScreenName = keyof StackParamList;

export type AppScreenProps<T extends ScreenName> = StackScreenProps<StackParamList, T>;

export default {
  Loading: 'Loading',
  SignUp: 'SignUp',
  SignIn: 'SignIn',
  Home: 'Home',
} as {[key in string]: ScreenName};
