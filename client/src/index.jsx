import ReactDOM from 'react-dom';
import React from 'react';
import App from './components/App';
import createStore from './store/createStore';
import initialState from './reducers/initialState';

const store = createStore(initialState);


ReactDOM.render(<App store={store} />, document.getElementById('root'));
