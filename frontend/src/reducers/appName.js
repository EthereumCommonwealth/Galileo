import { NAME as A } from '../constants/actionTypes';
import initialState from '../initialState';
import * as R from 'ramda';

export default function setName(state = initialState, { type, payload }) {
  switch (type) {
    case A.GET_NAME:
      return R.merge(state, payload);
    default:
      return state;
  }
};
