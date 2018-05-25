import React from 'react';
/* eslint import/no-extraneous-dependencies: ["error", {"devDependencies": true}] */
import { storiesOf } from '@storybook/react';
import 'normalize.css';
import User from '../src/components/User';
import Avatar from '../src/components/User/Avatar';

const sweetlie = {
  id: 820953585441771521,
  url: 'https://t.co/w1R2SQFN8r',
  name: 'Sweetie ♀',
  screenName: 'wowsweetlie',
  lang: 'fr',
  location: 'France',
  description: 'Anarchist, Coder, Woman Chrysalis, prefer she/her',
  backgroundImage: 'https://abs.twimg.com/images/themes/theme1/bg.png',
  image: 'https://pbs.twimg.com/profile_images/926709708555210752/PQFhN17n_bigger.jpg',
  banner: 'https://pbs.twimg.com/profile_banners/820953585441771521/1523533751',
  favorites: 337,
  followers: 91,
  friends: 101,
  statues: 1662,
  updatedAt: '2018-05-22T21:21:16.531685Z',
  createdAt: '2015-05-22T21:21:16.531685Z',
};

storiesOf('User', module)
  .add('avatar', () => <Avatar url={sweetlie.image} />)
  .add('simple user', () => <User user={sweetlie} />);
