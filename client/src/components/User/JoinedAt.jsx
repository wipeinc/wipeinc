import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';

const JoinP = styled.p`
  margin-right: 25px;
  text-align: right;
`;

const JoinedAt = ({ date }) => {
  const displayDate = new Date(date).toDateString();
  return (
    <JoinP>
      on twitter since {displayDate}
    </JoinP>
  );
};

JoinedAt.propTypes = {
  date: PropTypes.string.isRequired,
};


export default JoinedAt;
