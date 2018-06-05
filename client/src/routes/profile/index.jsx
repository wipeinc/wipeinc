import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import UserContainer from '../../components/User/Container';
import Timeline from '../../components/Timeline';

const Main = styled.main`
  display: flex;
`;

const StyledUser = styled(UserContainer)`
  flex: 1;
  max-width: 500px;
`;

const StyledTimelineDiv = styled.div`
  flex: 1;
  max-width: 500px;
`;

const Profile = ({ match }) => (
  <Main>
    <StyledUser screenName={match.params.screenName} />
    <StyledTimelineDiv>
      <Timeline
        screenName={match.params.screenName}
        theme="dark"
      />
    </StyledTimelineDiv>
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
