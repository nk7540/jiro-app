import React, {useCallback} from 'react';
import {gql, useQuery} from '@apollo/client';
import PostListItem from './PostListItem';
import {PostsQuery} from './__generated__/PostsQuery';
import ListWithError from 'components/ListWithError';
import Geolocation from '@react-native-community/geolocation';
import {geoFire, firebaseRef, distanceBetween} from 'services/firebase';
import {useAppDispatch, setDistances, setPost, useAppSelector} from 'services/store';
import {useFocusEffect} from '@react-navigation/native';
import {PostListStoresQuery} from './__generated__/PostListStoresQuery';
import {PostListUserQuery} from './__generated__/PostListUserQuery';

const POSTS_QUERY = gql`
  query PostsQuery {
    posts {
      ...post
    }
  }
  ${PostListItem.fragments.data}
`;

const POST_LIST_STORES_QUERY = gql`
  query PostListStoresQuery {
    stores {
      id
      name
      latitude
      longitude
    }
  }
`;

const POST_LIST_USER_QUERY = gql`
  query PostListUserQuery {
    currentUser {
      id
    }
  }
`;

interface Props {
  noUpdate: boolean;
}

const PostList = ({noUpdate}: Props) => {
  const post = useAppSelector(state => state.post);
  const dispatch = useAppDispatch();
  const {loading, error, data, refetch} = useQuery<PostsQuery>(POSTS_QUERY);
  const {data: storeData} = useQuery<PostListStoresQuery>(POST_LIST_STORES_QUERY);
  useQuery<PostListUserQuery>(POST_LIST_USER_QUERY, {
    onCompleted: currentUserData => {
      dispatch(setPost({key: 'userId', value: currentUserData.currentUser.id}));
    },
  });

  useFocusEffect(
    useCallback(() => {
      if (!storeData) {
        return;
      }
      const cancels = [] as (() => void)[];
      storeData.stores.forEach(store => {
        const query = geoFire.query({
          center: [store.latitude, store.longitude],
          radius: 0.1,
        });
        query.on('key_entered', (key: string) => {
          console.log('someone entered');
          if (!noUpdate && key === post.userId.toString() && post.store.id === 0) {
            dispatch(setPost({key: 'store', value: store}));
          }
        });
        query.on('key_exited', (key: string) => {
          console.log('someone exited');
          if (!noUpdate && key === post.userId.toString() && post.store.id === store.id) {
            dispatch(setPost({key: 'store', value: {id: 0, name: ''}}));
          }
        });
        cancels.push(query.cancel.bind(query));
      });

      return () => {
        cancels.forEach(cancel => cancel());
      };
    }, [storeData, dispatch, noUpdate, post.userId, post.store.id]),
  );

  const refreshList = useCallback(async () => {
    refetch();
    Geolocation.getCurrentPosition(
      position => {
        console.log(position);
        // This is just for an individual display. Leave the whole process of real-time suggesting to geofire.
        // This way you can easily integrate location data of other users. All comes in real-time.
        let newDistances = {} as {[storeId in number]: number};
        storeData &&
          storeData.stores.forEach(store => {
            const distance = distanceBetween(
              [position.coords.latitude, position.coords.longitude],
              [store.latitude, store.longitude],
            );
            newDistances = {...newDistances, [store.id]: distance};
          });
        dispatch(setDistances(newDistances));
        // Set current location to firebase realtime database.
        geoFire.set(post.userId.toString(), [position.coords.latitude, position.coords.longitude]).then(() => {
          // When the user disconnects from Firebase (e.g. closes the app, exits the browser),
          // remove their GeoFire entry
          firebaseRef.child(post.userId.toString()).onDisconnect().remove();
        });
      },
      error => console.log('Error', JSON.stringify(error)),
      {enableHighAccuracy: true, timeout: 5000, maximumAge: 1000},
    );
  }, [refetch, storeData, dispatch, post.userId]);

  return (
    <ListWithError loading={loading} error={error?.message} objects={data?.posts} refetch={refreshList}>
      {post => <PostListItem data={post} />}
    </ListWithError>
  );
};

export default PostList;
