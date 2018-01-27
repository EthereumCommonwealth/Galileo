import React from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import Home from './containers/Home';
import NotFound from './containers/NotFound';

export default () => {
 return (
   <BrowserRouter>
     <Switch>
       <Route exact path='/' component={Home} />
       <Route component={NotFound}/>
     </Switch>
   </BrowserRouter>
 );
}
