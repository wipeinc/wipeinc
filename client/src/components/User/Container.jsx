import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import fetchUser from '../../actions/userActions';

import User from './';

class Container extends React.Component {
  componentDidMount() {
    this.props.dispatch(fetchUser(this.props.screenName));
  }
  render() {
    const { user, loading, error } = this.props;
    if (error) {
      return <p>{error}</p>;
    }
    if (loading) {
      return <p>loading...</p>;
    }
    return <User user={user} />;
  }
}


Container.propTypes = {
  user: PropTypes.object,
  loading: PropTypes.bool,
  error: PropTypes.string,
  screenName: PropTypes.string.isRequired,
  dispatch: PropTypes.func.isRequired,
};

Container.defaultProps = {
  user: null,
  loading: true,
  error: null,
};

const mapStateToProps = ({ user }) => ({
  user: user.user,
  loading: user.loading,
  error: user.error,
});

export default connect(mapStateToProps)(Container);
