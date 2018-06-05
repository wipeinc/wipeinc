import React from 'react';
import PropTypes from 'prop-types';

const TwitterJS = ({ story }) => (
  <div>
    {story}
    <script async src="https://platform.twitter.com/widgets.js" charSet="utf-8" />
  </div>
);


TwitterJS.propTypes = {
  story: PropTypes.object.isRequired,
};


export default TwitterJS;

