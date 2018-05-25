import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';

const Definitions = styled.dl`
  display: grid;
`;

const Property = styled.dt`
  grid-column-start: 1;
`;

const Value = styled.dl`
  grid-column-start: 2;
  text-align: right;
`;

const UserStats = ({
  followers,
  friends,
  statues,
  favorites,
}) => (
  <Definitions>
    <Property>Tweets</Property>
    <Value>{statues}</Value>
    <Property>Favorites</Property>
    <Value>{favorites}</Value>
    <Property>Follwing</Property>
    <Value>{friends}</Value>
    <Property>Followers</Property>
    <Value>{followers}</Value>
  </Definitions>
);

UserStats.propTypes = {
  followers: PropTypes.number.isRequired,
  friends: PropTypes.number.isRequired,
  statues: PropTypes.number.isRequired,
  favorites: PropTypes.number.isRequired,
};

export default UserStats;
