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
  status: 'public',
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
    addTagId: (state, action: PayloadAction<number>) => {
      state.tagIds.push(action.payload);
    },
    removeTagId: (state, action: PayloadAction<number>) => {
      const index = state.tagIds.indexOf(action.payload);
      if (index > -1) {
        state.tagIds.splice(index, 1);
      }
    },
    resetPost: () => initialState,
  },
});

export const {setPost, increment, addTagId, removeTagId, resetPost} = slice.actions;
export default slice.reducer;
