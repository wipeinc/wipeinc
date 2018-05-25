import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import Avatar from './Avatar';


const Name = styled.p`
  color: white;
  font-size: 18px;
  margin-top: 10px;
  margin-bottom: 10px;
`;

const NameBox = styled.div`
  float: right;
  text-align: right;
`;

const IDBox = styled.div`
  padding: 10px;
`;

const StyledAvatar = styled(Avatar)`
  position: absolute;
`;

const Description = styled.p`
  margin-top: 14px;
  margin-bottom: 16px;
  margin-left: 50px;
  margin-right: 50px;
  color: white;
  text-align: center;
`;

const CardHeader = (props) => {
  const Box = styled.div`
    width: 500px;
    min-height: 100px;
    background-image: linear-gradient(
      rgba(0, 0, 0, 0.6),
      rgba(0, 0, 0, 0.6)
    ), url(${props.background});
    background-size: cover;
  `;

  const screenName = `@${props.screenName}`;

  return (
    <Box background={props.background} >
      <IDBox>
        <StyledAvatar url={props.avatar} />
        <NameBox>
          <Name>{props.name}</Name>
          <Name>{screenName}</Name>
        </NameBox>
        <Description>{props.description}</Description>
      </IDBox>
    </Box>
  );
};


CardHeader.propTypes = {
  avatar: PropTypes.string.isRequired,
  background: PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
  screenName: PropTypes.string.isRequired,
  description: PropTypes.string.isRequired,
};

export default CardHeader;
