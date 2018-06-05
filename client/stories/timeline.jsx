import 'normalize.css';
import React from 'react';
/* eslint-disable import/no-extraneous-dependencies */
import { storiesOf } from '@storybook/react';
/* eslint-enable */

import Timeline from '../src/components/Timeline';


storiesOf('Timeline', module)
  .add('sweetlie', () => (
    <Timeline
      name="Sweetie"
      width={400}
      screenName="wowsweetlie"
    />));
