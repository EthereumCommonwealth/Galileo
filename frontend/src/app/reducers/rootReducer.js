import { combineReducers } from 'redux';
import { routerReducer } from 'react-router-redux';

import mockReducer from './mockReducer';

const rootReducer = combineReducers(
  {
    routing: routerReducer,
    activeChain: mockReducer,
    chainList: mockReducer,
    latestBlocks: mockReducer,
    transactionsPerDay: mockReducer,
    marketInfo: mockReducer,
    transactionStatus: mockReducer,
    blockInfo: mockReducer,
    blockTransactions: mockReducer,
    addressSummary: mockReducer,
  });

export default rootReducer;
