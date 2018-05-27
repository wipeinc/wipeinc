import React from 'react';
import PropTypes from 'prop-types';
import UserContainer from '../../components/User/Container';

const Profile = ({ match }) => (
  <UserContainer screenName={match.params.screenName} />
);

Profile.propTypes = {
  match: PropTypes.shape({
    params: PropTypes.shape({
      screenName: PropTypes.string.isRequired,
    }).isRequired,
  }).isRequired,
};

export default Profile;
