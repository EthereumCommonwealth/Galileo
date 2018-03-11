import { createStore, applyMiddleware, compose  } from 'redux'
import reducers from './reducers/'
import initialState from './initialState'
import thunk from 'redux-thunk'

let composeEnhancers
let store
if (typeof window !== 'undefined') {
  composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose
  store = createStore(reducers, initialState, composeEnhancers(applyMiddleware(thunk)))
} else {
  store = createStore(reducers, initialState, applyMiddleware(thunk))
}

export default store
