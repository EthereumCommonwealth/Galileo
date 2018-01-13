import { createStore, applyMiddleware, compose  } from 'redux';
import reducers from './reducers/';
import initialState from './initialState';
import thunk from 'redux-thunk';

const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose
const store = createStore(reducers, initialState, composeEnhancers(applyMiddleware(thunk)));
export default store
