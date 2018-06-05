import 'normalize.css';
import React from 'react';
/* eslint-disable import/no-extraneous-dependencies */
import { storiesOf } from '@storybook/react';
/* eslint-enable */
import Provider from './Provider';
import user from '../fixtures/sweetlie.json';

import User from '../src/components/User';

const { twitterUser } = user;
storiesOf('User', module)
  .addDecorator(story => <Provider story={story()} />)
  .add('sweetlie', () => (
    <User
      user={twitterUser}
      screenName={twitterUser.screen_name}
      loading={false}
      error={null}
    />));
