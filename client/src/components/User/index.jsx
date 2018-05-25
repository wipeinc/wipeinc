import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import CardHeader from './CardHeader';
import UserStats from './UserStats';

const Card = styled.div`
  width: 500px;
`;

const User = ({ user }) => (
  <Card>
    <div className="card-content">
      <CardHeader
        background={user.banner}
        avatar={user.image}
        name={user.name}
        screenName={user.screenName}
        description={user.description}
      />
      <UserStats
        followers={user.followers}
        friends={user.friends}
        statues={user.statues}
        favorites={user.favorites}
      />
    </div>
  </Card>
);

User.propTypes = {
  user: PropTypes.shape({
    name: PropTypes.string.isRequired,
    screenName: PropTypes.string.isRequired,
    image: PropTypes.string.isRequired,
    banner: PropTypes.string.isRequired,
    description: PropTypes.string.isRequired,
    followers: PropTypes.number.isRequired,
    friends: PropTypes.number.isRequired,
    statues: PropTypes.number.isRequired,
    favorites: PropTypes.number.isRequired,
  }).isRequired,
};


export default User;
