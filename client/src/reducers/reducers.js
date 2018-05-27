import { combineReducers } from 'redux';
import userReducer from './userReducer';

const makeRootReducer = asyncReducers => (
  combineReducers({
    user: userReducer,
    ...asyncReducers,
  })
);


export const injectReducer = (store, { key, reducer }) => {
  if (Object.hasOwnProperty.call(store.asyncReducers, key)) return;

  /* eslint-disable no-param-reassign */
  store.asyncReducers[key] = reducer;
  /* estlint-enable */
  store.replaceReducer(makeRootReducer(store.asyncReducers));
};

export default makeRootReducer;
