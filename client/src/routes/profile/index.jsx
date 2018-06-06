import React from 'react';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';
import styled from 'styled-components';
import UserContainer from '../../components/User/Container';
import Timeline from '../../components/Timeline';

const Main = styled.main`
  display: flex;
  background-color: #1e1d23;
`;

const UserDiv = styled.div`
  flex: 1;
  max-width: 520px;
`;

const TimelineDiv = styled.div`
  flex: 1;
  max-width: 520px;
  margin: 7px;
`;

const Button = styled(Link)`
  display: inline-block;
  background-color: #25252f;
  color: white;
  vertical-align: middle;
  text-align: center;
  padding: 0.5rem 1rem;
  width: calc(100% - 2rem - 14px);
  margin: 7px;
  border-radius: 5px;
`;

const Profile = ({ match }) => (
  <Main>
    <UserDiv>
      <UserContainer screenName={match.params.screenName} />
      <Button to="/profile/{screenName}/analyze">
        Analyze Profile
      </Button>
    </UserDiv>
    <TimelineDiv>
      <Timeline
        screenName={match.params.screenName}
        theme="dark"
      />
    </TimelineDiv>
  </Main>
);

Profile.propTypes = {
  match: PropTypes.shape({
    params: PropTypes.shape({
      screenName: PropTypes.string.isRequired,
    }).isRequired,
  }).isRequired,
};

export default Profile;
