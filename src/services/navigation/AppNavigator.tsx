import React, {useEffect, useContext, FunctionComponent, useState} from 'react';
import {NavigationContainer} from '@react-navigation/native';
import {createStackNavigator} from '@react-navigation/stack';
import {StackParamList, ScreenName} from './ScreenNames';
import NavigationService from './NavigationService';
import {Status, StatusContext} from 'services/context/StatusProvider';
import Loading from 'features/Loading';
import SignIn from 'features/SignIn';
import SignUp from 'features/SignUp';
import Home from 'features/Home';
import CreateImages from 'features/CreateImages';
import CreatePost from 'features/CreatePost';
import {Linking, AsyncStorage} from 'react-native';

// see: https://reactnavigation.org/docs/state-persistence/
const PERSISTENCE_KEY = 'NAVIGATION_STATE';

const Stack = createStackNavigator<StackParamList>();

const loadingScreen = {
  Loading: Loading,
};

const authScreens = {
  SignIn: SignIn,
  SignUp: SignUp,
};

const userScreens = {
  Home: Home,
  CreateImages: CreateImages,
  CreatePost: CreatePost,
};

function switchByStatus(status: Status): JSX.Element[] {
  let screens = {} as {[name in ScreenName]: FunctionComponent};
  switch (status) {
    case Status.LOADING:
      screens = Object.assign(screens, loadingScreen);
      break;
    case Status.UN_AUTHORIZED:
      screens = Object.assign(screens, authScreens);
      break;
    case Status.AUTHORIZED:
      screens = Object.assign(screens, userScreens);
  }
  return Object.entries(screens).map(([name, component]) => (
    <Stack.Screen key={name} name={name as ScreenName} component={component} />
  ));
}

export default function AppNavigator() {
  const {status} = useContext(StatusContext);
  const [isReady, setIsReady] = useState(false);
  const [initialState, setInitialState] = useState();

  useEffect(() => {
    const restoreState = async () => {
      try {
        const initialUrl = await Linking.getInitialURL();

        if (initialUrl == null) {
          // Only restore state if there's no deep link
          const savedStateString = await AsyncStorage.getItem(PERSISTENCE_KEY);
          const state = savedStateString ? JSON.parse(savedStateString) : undefined;

          if (state !== undefined) {
            setInitialState(state);
          }
        }
      } finally {
        setIsReady(true);
      }
    };

    if (!isReady) {
      restoreState();
    }

    return () => {
      Object.assign(NavigationService.isReadyRef, {current: false});
    };
  }, [isReady]);

  if (!isReady) {
    return null;
  }

  return (
    <NavigationContainer
      initialState={initialState}
      onStateChange={state => AsyncStorage.setItem(PERSISTENCE_KEY, JSON.stringify(state))}
      ref={NavigationService.navigationRef}
      onReady={() => Object.assign(NavigationService.isReadyRef, {current: true})}>
      <Stack.Navigator screenOptions={{headerShown: false}}>{switchByStatus(status)}</Stack.Navigator>
    </NavigationContainer>
  );
}
