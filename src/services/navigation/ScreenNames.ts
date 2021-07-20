import {StackScreenProps} from '@react-navigation/stack';

export type StackParamList = {
  Loading: undefined;
  SignUp: undefined;
  SignIn: undefined;
  Home: undefined;
  CreateImages: undefined;
};

export type ScreenName = keyof StackParamList;

export type AppScreenProps<T extends ScreenName> = StackScreenProps<StackParamList, T>;

export default {
  Loading: 'Loading',
  SignUp: 'SignUp',
  SignIn: 'SignIn',
  Home: 'Home',
  CreateImages: 'CreateImages',
} as {[key in string]: ScreenName};
