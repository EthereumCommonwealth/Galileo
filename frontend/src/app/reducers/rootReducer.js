import { combineReducers } from 'redux';
import { routerReducer } from 'react-router-redux';

import authReducer from './authReducer';

const rootReducer = combineReducers(
  {
    routing: routerReducer,
    authenticated: authReducer,
  });

export default rootReducer;
