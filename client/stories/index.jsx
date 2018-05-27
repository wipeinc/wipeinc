import 'normalize.css';
import React from 'react';
/* eslint import/no-extraneous-dependencies: ["error", {"devDependencies": true}] */
import { storiesOf } from '@storybook/react';
import { Provider } from './Provider';


import User from '../src/components/User';


const sweetlie = {
  "id": 820953585441771521,
  "url": "https://t.co/w1R2SQFN8r",
  "name": "Sweetie ♀",
  "screenName": "wowsweetlie",
  "location": "France",
  "lang": "fr",
  "description": "Anarchist, Coder, Woman Chrysalis, prefer she/her",
  "backgroundImage": "https://abs.twimg.com/images/themes/theme1/bg.png",
  "image": "https://pbs.twimg.com/profile_images/926709708555210752/PQFhN17n_normal.jpg",
  "banner": "https://pbs.twimg.com/profile_banners/820953585441771521/1523533751",
  "statuses": 1702,
  "favorites": 375,
  "followers": 92,
  "friends": 102,
  "createdAt": "2017-01-16T11:19:22Z",
  "updatedAt": "2018-05-25T11:39:22.045269Z"
}


storiesOf('User', module)
  .addDecorator(story => <Provider story={story()} />)
  .add('sweetlie', () => (
    <User
      user={sweetlie}
      screenName={sweetlie.screenName}
      loading={false}
      error={null}
    />));
