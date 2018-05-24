import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';

const Img = styled.img`
  border-radius: 5px;
`;

const Avatar = ({ url, size }) => (
  <div>
    <Img src={url} width={size} height={size} alt="avatar" />
  </div>
);

Avatar.propTypes = {
  url: PropTypes.string.isRequired,
  size: PropTypes.number,
};

Avatar.defaultProps = {
  size: 72,
};

export default Avatar;
