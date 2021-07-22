export {default as store} from './store';
export type {AppDispatch, RootState, AsyncThunkConfig} from './store';
export {useAppDispatch, useAppSelector} from './useAppStore';

export {setPost, increment, addTagId, removeTagId, resetPost} from './post';
