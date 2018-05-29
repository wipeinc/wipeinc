import 'normalize.css';
import React from 'react';
/* eslint import/no-extraneous-dependencies: ["error", {"devDependencies": true}] */
import { storiesOf } from '@storybook/react';
import { Provider } from './Provider';
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
