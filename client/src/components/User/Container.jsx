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
    console.log("trigerring render");
    console.dir(this.props);
    const { user, loading, error } = this.props;
    if (error) {
      return <p>{error}</p>;
    }
    if (loading || !user) {
      return <p>loading...</p>;
    }
    return <User user={user} />;
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
  const { user = {}, loading = true, error = '' } = state.user.toJS();
  console.log('mapsToProps');
  console.dir(state.user);
  console.dir(user);
  return ({
    user,
    loading,
    error,
  });
};

export default connect(mapStateToProps, mapDispatchToProps)(Container);
