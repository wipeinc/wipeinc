import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';

const Img = styled.img`
  border-radius: 50%;
  overflow: hidden;
`;

const Avatar = ({ url, size }) => {
  const biggerURL = url.replace('normal', 'bigger');
  return (
    <Img src={biggerURL} width={size} height={size} alt="avatar" />
  );
};

Avatar.propTypes = {
  url: PropTypes.string.isRequired,
  size: PropTypes.number,
};

Avatar.defaultProps = {
  size: 72,
};

export default Avatar;
