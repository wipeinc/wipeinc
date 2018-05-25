import ReactDOM from 'react-dom';
import React from 'react';
import App from './components/App';
import createStore from './store/createStore';

/* eslint-disable no-underscore-dangle */
const store = createStore(window.__INITIAL_STATE__);
/* estlin-enable */

ReactDOM.render(<App store={store} />, document.getElementById('root'));
