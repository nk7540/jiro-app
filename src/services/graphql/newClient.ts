import {ApolloClient, InMemoryCache, from} from '@apollo/client';
import {setContext} from '@apollo/client/link/context';
import {onError} from '@apollo/client/link/error';
import createUploadLink from 'apollo-upload-client/public/createUploadLink.js';
import auth from '@react-native-firebase/auth';
import {signOutFromFirebase} from 'services/firebase';
import {TSetStatus, Status} from 'services/context/StatusProvider';

const uploadLink = createUploadLink({uri: 'http://localhost:8080/query'});

const authLink = setContext(async (_, {headers}) => {
  // get the authentication token from local storage if it exists
  const token = await auth().currentUser?.getIdToken();
  console.log(token);
  // return the headers to the context so httpLink can read them
  return {
    headers: {
      ...headers,
      authorization: token ? `Bearer ${token}` : '',
    },
  };
});

const newErrorLink = (setStatus: TSetStatus) =>
  onError(({graphQLErrors, networkError}) => {
    if (graphQLErrors) {
      graphQLErrors.forEach(({message, locations, path}) => {
        console.log(`[GraphQL error]: Message: ${message}, Location: ${locations}, Path: ${path}`);
        if (message === 'access denied') {
          signOutFromFirebase();
          setStatus(Status.UN_AUTHORIZED);
        }
      });
    }

    if (networkError) {
      console.log(JSON.stringify(networkError));
      console.log(`[Network error]: ${networkError}`);
    }
  });

const newClient = (setStatus: TSetStatus) =>
  new ApolloClient({
    link: from([newErrorLink(setStatus), authLink, uploadLink]),
    cache: new InMemoryCache(),
    connectToDevTools: true,
  });

export default newClient;
