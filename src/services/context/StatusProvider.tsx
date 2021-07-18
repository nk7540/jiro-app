import React, {ReactNode, useState, createContext, Dispatch, SetStateAction} from 'react';

export enum Status {
  LOADING = 'loading',
  UN_AUTHORIZED = 'unAuthorized',
  AUTHORIZED = 'authorized',
}

export type TSetStatus = Dispatch<SetStateAction<Status>>;

interface IStatusContext {
  status: Status;
  setStatus: TSetStatus;
}

export const StatusContext = createContext({} as IStatusContext);

const StatusProvider = (props: {children: ReactNode}) => {
  const [status, setStatus] = useState<Status>(Status.LOADING);

  return <StatusContext.Provider value={{status, setStatus}}>{props.children}</StatusContext.Provider>;
};

export default StatusProvider;
