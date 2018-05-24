import React from 'react';
import PropTypes from 'prop-types';
import Avatar from './Avatar';

const User = ({ user }) => (
  <div className="card">
    <div className="card-content">
      <Avatar url={user.image} />
      <div className="media-content">
        <p className="title is-4">{user.name}</p>
        <p className="subtitle is-6">{user.screenName}</p>
      </div>
      <div className="field is-grouped">
        <div className="control">
          <div className="tags">
            <span className="tag">F</span>
            <span className="tag">12</span>
          </div>
        </div>
      </div>
    </div>
  </div>
);

User.propTypes = {
  user: PropTypes.shape({
    name: PropTypes.string.isRequired,
    screenName: PropTypes.string.isRequired,
    image: PropTypes.string.isRequired,
  }).isRequired,
};


export default User;
