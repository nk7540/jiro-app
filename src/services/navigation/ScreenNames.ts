import {StackScreenProps} from '@react-navigation/stack';

export type StackParamList = {
  Loading: undefined;
  SignUp: undefined;
  SignIn: undefined;
  Main: undefined;
  Home: undefined;
  StoreList: undefined;
  CreateImages: undefined;
  CreatePost: undefined;
};

export type ScreenName = keyof StackParamList;

export type AppScreenProps<T extends ScreenName> = StackScreenProps<StackParamList, T>;

export default {
  Loading: 'Loading',
  SignUp: 'SignUp',
  SignIn: 'SignIn',
  Main: 'Main',
  Home: 'Home',
  StoreList: 'StoreList',
  CreateImages: 'CreateImages',
  CreatePost: 'CreatePost',
} as {[key in string]: ScreenName};
