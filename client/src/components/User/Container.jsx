import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import fetchUser from '../../actions/userActions';

import User from './';

class Container extends React.Component {
  componentDidMount() {
    this.props.fetchUser(this.props.screenName);
  }
  render() {
    const { user, loading, error } = this.props;
    if (error) {
      return <p>{error}</p>;
    }
    if (loading || !Object.keys(user).length) {
      return <p>loading...</p>;
    }
    return <User user={user.twitterUser} />;
  }
}


Container.propTypes = {
  user: PropTypes.object.isRequired,
  loading: PropTypes.bool.isRequired,
  error: PropTypes.string.isRequired,
  screenName: PropTypes.string.isRequired,
  fetchUser: PropTypes.func.isRequired,
};

const mapDispatchToProps = {
  fetchUser,
};

const mapStateToProps = (state) => {
  const { user, loading, error } = state.user.toJS();
  return ({
    user,
    loading,
    error,
  });
};

export default connect(mapStateToProps, mapDispatchToProps)(Container);
