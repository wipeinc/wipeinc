import React from 'react';
import PropTypes from 'prop-types';

const Profile = ({ match }) => (
  <h1>loading {match.params.screenName} profile</h1>
);

Profile.propTypes = {
  match: PropTypes.shape({
    screenName: PropTypes.shape({
      screenName: PropTypes.string.isRequired,
    }).isRequired,
  }).isRequired,
};


export default Profile;
