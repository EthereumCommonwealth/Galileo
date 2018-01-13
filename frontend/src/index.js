import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import App from './containers/App';
import store from './store';
import registerServiceWorker from './registerServiceWorker';
import './client/styles/index.css';
import 'normalize.css/normalize.css';

ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('root')
);

registerServiceWorker();
