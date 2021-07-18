import React from 'react';
import {NavigationContainerRef} from '@react-navigation/native';

const isReadyRef = React.createRef<NavigationContainerRef>();
const navigationRef = React.createRef<NavigationContainerRef>();

const navigate = (name: string, params: any = {}) => {
  if (isReadyRef.current && navigationRef.current) {
    navigationRef.current.navigate(name, params);
  }
};

export default {
  isReadyRef,
  navigationRef,
  navigate,
};
