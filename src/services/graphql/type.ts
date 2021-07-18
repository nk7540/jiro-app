import React from 'react';
import {DocumentNode} from 'graphql';

export interface FragC<FragmentProps, Props = {}> extends React.FC<Props & FragmentProps> {
  fragments: Record<keyof FragmentProps, DocumentNode>;
}
