import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import Profile from './pages/profile/Profile';

const App = () => (
  <Router>
    <Route path="/profile/:screenName" component={Profile} />
  </Router>
);


ReactDOM.render(<App />, document.getElementById('root'));
