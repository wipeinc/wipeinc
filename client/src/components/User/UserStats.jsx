import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';

const Definitions = styled.dl`
  display: grid;
  margin: 20px 10px;
`;

const Property = styled.dt`
  grid-column-start: 1;
  padding: 10px 15px;
  &:not(:last-of-type) {
    border-bottom: 1px solid #444;
  }
`;

const Value = styled.dl`
  grid-column-start: 2;
  text-align: right;
  padding: 10px 15px;
  &:not(:last-of-type) {
    border-bottom: 1px solid #444;
  }
`;

const UserStats = ({
  followers,
  friends,
  statuses,
  favorites,
}) => (
  <Definitions>
    <Property>Tweets</Property>
    <Value>{statuses}</Value>
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
  statuses: PropTypes.number.isRequired,
  favorites: PropTypes.number.isRequired,
};

export default UserStats;
