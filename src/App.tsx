import React from 'react';
import {Provider} from 'react-redux';
import {ThemeProvider} from 'react-native-elements';
import {ApolloProvider} from '@apollo/client';

import {AppNavigator} from 'services/navigation';
import {theme} from 'services/theme';
import {newClient} from 'services/graphql';
import StatusProvider, {StatusContext} from 'services/context/StatusProvider';
import {store} from 'services/store';

const App = () => {
  return (
    <Provider store={store}>
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
    </Provider>
  );
};

export default App;
