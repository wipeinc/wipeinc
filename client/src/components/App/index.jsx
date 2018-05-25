import React from 'react';
import { Provider } from 'react-redux';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import PropTypes from 'prop-types';
import Profile from '../../routes/profile';

const App = ({ store }) => (
  <Provider store={store}>
    <Router>
      <Route path="/profile/:screenName" component={Profile} />
    </Router>
  </Provider>
);

App.propTypes = {
  /* eslint-disable */
  store: PropTypes.object.isRequired,
  /* eslint-enable */
};

export default App;
