import {createSlice, PayloadAction} from '@reduxjs/toolkit';

type DistancesState = {[storeId in number]: number};
const initialState = {} as DistancesState;

const slice = createSlice({
  name: 'distances',
  initialState,
  reducers: {
    setDistances: (_, action: PayloadAction<DistancesState>) => action.payload,
  },
});

export const {setDistances} = slice.actions;
export default slice.reducer;
