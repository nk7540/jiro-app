import {configureStore} from '@reduxjs/toolkit';
import postReducer from './post';

const store = configureStore({
  reducer: {
    post: postReducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
export interface AsyncThunkConfig<E = any> {
  state: RootState;
  rejectValue: E;
}

export default store;
