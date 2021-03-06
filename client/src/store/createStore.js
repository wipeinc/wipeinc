import {
  applyMiddleware,
  compose,
  createStore as createReduxStore,
} from 'redux';
import thunk from 'redux-thunk';
import makeRootReducer from '../reducers/reducers';


const createStore = (initialState = {}) => {
  const middleware = [thunk];

  const enhancers = [];
  /* eslint-disable no-underscore-dangle */
  const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;
  /* eslint-enable */


  const store = createReduxStore(
    makeRootReducer(),
    initialState,
    composeEnhancers(
      applyMiddleware(...middleware),
      ...enhancers,
    ),
  );
  store.asyncReducers = {};
  if (module.hot) {
    module.hot.accept('../reducers/reducers', () => {
      /* eslint-disable global-require */
      const reducers = require('../reducers/reducers').default;
      /* eslint-enable */
      store.replaceReducer(reducers(store.asyncReducers));
    });
  }

  return store;
};

export default createStore;
