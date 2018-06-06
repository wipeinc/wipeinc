import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import FontAwesomeIcon from '@fortawesome/react-fontawesome';
import faTwitter from '@fortawesome/fontawesome-free-brands/faTwitter';
import Avatar from './Avatar';


const Name = styled.p`
  font-size: 18px;
  margin-top: 7px;
  margin-bottom: 7px;
  text-shadow:0 2px 1px rgba(0,0,0,.6);
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

const TwitterLink = styled.a`
  font-size: 18px;
  margin-top: 0px;
  margin-bottom: 12px;
  text-shadow:0 2px 1px rgba(0,0,0,.6);
  color: #1da1f2;
  text-decoration: none;
`;

const Description = styled.p`
  margin-top: 20px;
  margin-bottom: 16px;
  margin-left: 50px;
  margin-right: 50px;
  text-align: center;
  text-shadow:0 2px 1px rgba(0,0,0,.6);
`;

const CardHeader = (props) => {
  const Box = styled.div`
    width: 500px;
    min-height: 100px;
    border-radius: 5px;
    background-image: linear-gradient(
      rgba(0, 0, 0, 0.6),
      rgba(0, 0, 0, 0.6)
    ), url(${props.background});
    background-size: cover;
  `;

  const screenName = `@${props.screenName}`;
  const url = `https://twitter.com/${props.screenName}`;


  return (
    <Box background={props.background} >
      <IDBox>
        <StyledAvatar url={props.avatar} />
        <NameBox>
          <TwitterLink href={url}>
            profile <FontAwesomeIcon icon={faTwitter} size="xs" />
          </TwitterLink>
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
