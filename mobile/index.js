import React from 'react';
import { AppRegistry } from 'react-native';
import { createStore, applyMiddleware, compose } from 'redux';
import thunk from 'redux-thunk';
import { Provider } from 'react-redux';
import initialState from './initialState';
import App from './containers/App';
import rootReducer from './reducers/rootReducer';

const store = createStore(rootReducer, initialState, compose(applyMiddleware(thunk)));
const GalileoApp = () => (
  <Provider store={store}>
    <App />
  </Provider>
);

AppRegistry.registerComponent('mobile', () => GalileoApp);
