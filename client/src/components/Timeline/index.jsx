import React from 'react';
import PropTypes from 'prop-types';

class Timeline extends React.Component {
  constructor(props) {
    super(props);
    this.state = ({ initialized: false });
  }

  componentDidMount() {
    if (this.state.initialized) {
      return;
    }

    if (typeof twttr === 'undefined') {
      const twittertimeline = this.node;
      const twitterscript = document.createElement('script');
      twitterscript.src = 'https://platform.twitter.com/widgets.js';
      twitterscript.async = true;
      twitterscript.id = 'twitter-wjs';
      twittertimeline.parentNode.appendChild(twitterscript);
    } else {
      twttr.widgets.load();
    }

    this.initialized();
  }

  initialized() {
    this.setState({ initialized: true });
  }

  render() {
    const {
      name,
      screenName,
      width,
      height,
      limit,
    } = this.props;
    const loading = `loading Tweets by ${name}`;
    const link = `https://twitter.com/${screenName}`;
    return (
      <a
        ref={node => this.node = node}
        className="twitter-timeline"
        href={link}
        data-tweet-limit={limit}
        data-width={width}
        data-height={height}
      >
        {loading}
      </a>
    );
  }
}

Timeline.propTypes = {
  screenName: PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
  width: PropTypes.number,
  height: PropTypes.number,
  limit: PropTypes.number,
};

Timeline.defaultProps = {
  width: null,
  height: null,
  limit: null,
};

export default Timeline;
