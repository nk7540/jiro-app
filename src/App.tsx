import React from 'react';
import {Provider} from 'react-redux';
import {ThemeProvider} from 'react-native-elements';
import {ApolloProvider} from '@apollo/client';

import {AppNavigator} from 'services/navigation';
import {theme} from 'services/theme';
import {newClient} from 'services/graphql';
import StatusProvider, {StatusContext} from 'services/context/StatusProvider';
import {store, persistor} from 'services/store';
import {PersistGate} from 'redux-persist/integration/react';

const App = () => {
  return (
    <Provider store={store}>
      <PersistGate loading={null} persistor={persistor}>
        <ThemeProvider theme={theme}>
          <StatusProvider>
            <StatusContext.Consumer>
              {({setStatus}) => (
                <ApolloProvider client={newClient(setStatus)}>
                  <AppNavigator />
                </ApolloProvider>
              )}
            </StatusContext.Consumer>
          </StatusProvider>
        </ThemeProvider>
      </PersistGate>
    </Provider>
  );
};

export default App;
