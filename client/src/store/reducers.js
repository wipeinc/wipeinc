import { combineReducers } from 'redux';

const makeRootReducer = asyncReducers => (
  combineReducers({
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
