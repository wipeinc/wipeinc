import { configure } from '@storybook/react';

function loadStories() {
  require('../stories/index.jsx');
  require('../stories/timeline.jsx');
  // You can require as many stories as you need.
}

configure(loadStories, module);
