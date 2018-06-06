import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';

import CardHeader from './CardHeader';
import UserStats from './UserStats';
import JoinedAt from './JoinedAt';

const Card = styled.div`
  width: 500px;
  background-color: #1e1d23;
  color: white;
  padding: 10px;
`;

const CardContent = styled.div`
  background-color: #25252f;
  margin-top: 10px;
  padding-top:  10px;
  padding-bottom: 10px;
  border-radius: 5px;
`;

const User = ({ user }) => (
  <Card>
    <div className="card-content">
      <CardHeader
        background={user.profile_banner_url}
        avatar={user.profile_image_url_https}
        name={user.name}
        screenName={user.screen_name}
        description={user.description}
      />
      <CardContent>
        <UserStats
          followers={user.followers_count}
          friends={user.friends_count}
          statuses={user.statuses_count}
          favorites={user.favourites_count}
        />
        <JoinedAt date={user.created_at} />
      </CardContent>
    </div>
  </Card>
);

User.propTypes = {
  user: PropTypes.shape({
    name: PropTypes.string.isRequired,
    screen_name: PropTypes.string.isRequired,
    profile_image_url_https: PropTypes.string.isRequired,
    profile_banner_url: PropTypes.string.isRequired,
    description: PropTypes.string.isRequired,
    followers_count: PropTypes.number.isRequired,
    friends_count: PropTypes.number.isRequired,
    statuses_count: PropTypes.number.isRequired,
    favourites_count: PropTypes.number.isRequired,
    created_at: PropTypes.string.isRequired,
  }).isRequired,
};

export default User;
