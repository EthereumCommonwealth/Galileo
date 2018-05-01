import React from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import Home from '../containers/Home';
import Chain from '../containers/Chain';
import Transaction from '../containers/Transaction';
import Block from '../containers/Block';
import BlockTransactions from '../containers/BlockTransactions';
import Notfound from '../containers/NotFound';

const Routes = (
  <BrowserRouter>
    <Switch>
      <Route exact path='/' component={Home} />
      <Route exact path='/chain/:chain' component={Chain} />
      <Route exact path='/chain/:chain/transaction/:hash' component={Transaction} />
      <Route exact path='/chain/:chain/block/:hash' component={Block} />
      <Route exact path='/chain/:chain/block/:hash/transactions' component={BlockTransactions} />
      <Route component={Notfound} />
    </Switch>
  </BrowserRouter>
);

export default Routes;
