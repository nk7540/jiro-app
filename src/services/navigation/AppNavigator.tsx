import React, {useEffect, useContext, FunctionComponent} from 'react';
import {NavigationContainer} from '@react-navigation/native';
import {createStackNavigator} from '@react-navigation/stack';
import {StackParamList, ScreenName} from './ScreenNames';
import NavigationService from './NavigationService';
import {Status, StatusContext} from 'services/context/StatusProvider';
import Loading from 'features/Loading';
import SignIn from 'features/SignIn';
import SignUp from 'features/SignUp';
import Home from 'features/Home';

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

  useEffect(() => {
    return () => {
      Object.assign(NavigationService.isReadyRef, {current: false});
    };
  }, []);

  return (
    <NavigationContainer
      ref={NavigationService.navigationRef}
      onReady={() => Object.assign(NavigationService.isReadyRef, {current: true})}>
      <Stack.Navigator screenOptions={{headerShown: false}}>{switchByStatus(status)}</Stack.Navigator>
    </NavigationContainer>
  );
}
