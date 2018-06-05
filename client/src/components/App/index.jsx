import React from 'react';
import { Provider } from 'react-redux';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import PropTypes from 'prop-types';
import Profile from '../../routes/profile';
import Home from '../../routes/home';

import '../../style/index.scss';


const App = ({ store }) => (
  <Provider store={store}>
    <Router>
      <Switch>
        <Route exact path="/" component={Home} />
        <Route path="/profile/:screenName" component={Profile} />
      </Switch>
    </Router>
  </Provider>
);

App.propTypes = {
  /* eslint-disable */
  store: PropTypes.object.isRequired,
  /* eslint-enable */
};

export default App;
