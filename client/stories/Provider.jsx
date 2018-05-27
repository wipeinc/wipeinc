import React from 'react';
import { Provider as ReduxProvider } from 'react-redux';
import createStore from '../src/store/createStore';

/* eslint-disable no-underscore-dangle */
const store = createStore(window.__INITIAL_STATE__);
/* estlin-enable */

const Provider = ({ story }) => (
  <ReduxProvider store={store}>
    {story}
  </ReduxProvider>
);

export default Provider;
