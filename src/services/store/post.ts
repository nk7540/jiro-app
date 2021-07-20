import {ReactNativeFile} from 'apollo-upload-client';
import {createSlice, PayloadAction} from '@reduxjs/toolkit';
import {store} from 'features/Home/__generated__/store';

const initialState = {
  userId: 0,
  store: {id: 0, name: ''},
  waitingFor: 0,
  images: [] as ReactNativeFile[],
  consumingFor: 0,
  comment: '',
  tagIds: [] as number[],
};

type PostState = typeof initialState;

const slice = createSlice({
  name: 'post',
  initialState,
  reducers: {
    setPost: <T extends keyof PostState>(state: PostState, action: PayloadAction<{key: T; value: PostState[T]}>) => {
      state[action.payload.key] = action.payload.value;
    },
    increment: (state, action: PayloadAction<'waitingFor' | 'consumingFor'>) => {
      state[action.payload] += 1;
    },
    resetPost: () => initialState,
  },
});

export const {setPost, increment, resetPost} = slice.actions;
export default slice.reducer;
