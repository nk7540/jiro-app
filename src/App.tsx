import React from 'react';
import {ThemeProvider} from 'react-native-elements';
import {ApolloProvider} from '@apollo/client';

import {AppNavigator} from 'services/navigation';
import {theme} from 'services/theme';
import {newClient} from 'services/graphql';
import StatusProvider, {StatusContext} from 'services/context/StatusProvider';

const App = () => {
  return (
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
  );
};

export default App;
