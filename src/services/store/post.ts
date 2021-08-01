import {createSlice, PayloadAction} from '@reduxjs/toolkit';

const initialState = {
  userId: 0,
  store: {id: 0, name: ''},
  waitingSince: 0,
  waitingFor: 0,
  images: [] as {uri: string; name: string; type: string}[],
  consumingSince: 0,
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
    updateWaitingFor: (state, action: PayloadAction<number>) => {
      state.waitingFor = Math.round((action.payload - state.waitingSince) / 1000);
    },
    resetWaiting: state => {
      state.waitingFor = 0;
      state.waitingSince = 0;
    },
    updateConsumingFor: (state, action: PayloadAction<number>) => {
      state.consumingFor = Math.round((action.payload - state.consumingSince) / 1000);
    },
    resetConsuming: state => {
      state.consumingFor = 0;
      state.consumingSince = 0;
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

export const {
  setPost,
  updateWaitingFor,
  resetWaiting,
  updateConsumingFor,
  resetConsuming,
  addTagId,
  removeTagId,
  resetPost,
} = slice.actions;
export default slice.reducer;
