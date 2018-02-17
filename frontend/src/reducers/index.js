import { combineReducers } from 'redux';

import appName from './appName';

const rootReducer = combineReducers(
  {
    name: appName,
  });

export default rootReducer;
