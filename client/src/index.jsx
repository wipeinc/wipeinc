import React from 'react';
import ReactDOM from 'react-dom';
import PropTypes from 'prop-types';
import { BrowserRouter as Router, Route } from 'react-router-dom';


const Profile = ({ match }) => (
  <h1>loading {match.params.screenName} profile</h1>
);

Profile.propTypes = {
  match: PropTypes.object.isRequired,
};

const App = () => (
  <Router>
    <Route path="/profile/:screenName" component={Profile} />
  </Router>
);


ReactDOM.render(<App />, document.getElementById('root'));
