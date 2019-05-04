import * as React from 'react';
import * as ReactDOM from 'react-dom';
import { HashRouter as Router, Switch, Route } from 'react-router-dom';
import Home from './Home';
import Settings from './Settings';

function NoMatch() {
  return (
    <div></div>
  )
}

ReactDOM.render(
  <Router>
    <Switch>
      <Route exact path="/" component={Home} />
      <Route exact path="/settings" component={Settings} />
      <Route component={NoMatch} />
    </Switch>
  </Router>,
  document.getElementById('app'),
);
