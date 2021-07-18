import React, {useContext} from 'react';
import {ActivityIndicator, StyleSheet, View} from 'react-native';
import auth from '@react-native-firebase/auth';
import {StatusContext, Status} from 'services/context/StatusProvider';

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
});

const Loading = () => {
  const {setStatus} = useContext(StatusContext);
  async function loading() {
    try {
      const user = await auth().currentUser;

      if (user) {
        // Navigate to Home
        setStatus(Status.AUTHORIZED);
      } else {
        // Navigate to SignIn page
        setStatus(Status.UN_AUTHORIZED);
      }
    } catch (err) {}
  }

  React.useEffect(() => {
    loading();
  });

  return (
    <View style={styles.container}>
      <ActivityIndicator size="large" />
    </View>
  );
};

export default Loading;
