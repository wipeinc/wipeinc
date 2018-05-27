import fetchMock from 'fetch-mock';
import configureMockStore from 'redux-mock-store';
import thunk from 'redux-thunk';

import fetchUser, * as userActions from '../../src/actions/userActions';
import * as types from '../../src/actions/actionTypes';
import user from '../../fixtures/sweetlie.json';

const middlewares = [thunk];
const mockStore = configureMockStore(middlewares);


describe('userActions', () => {
  afterEach(() => {
    fetchMock.reset();
    fetchMock.restore();
  });

  it('should set loading to true when downloading start', () => {
    const screenName = '@screen_name';
    const expectedAction = {
      type: types.FETCH_USER_BEGIN,
      payload: { screenName },
    };
    expect(userActions.fetchUserBegin(screenName)).toEqual(expectedAction);
  });

  it('fetch user and return FETCH_USER_SUCCESS when its done', () => {
    fetchMock.getOnce(`${__API_BASE_URL__}/user/${user.screenName}`, { body: user, headers: { 'content-type': 'application/json' } });
    const expectedActions = [
      { type: types.FETCH_USER_BEGIN, payload: { screenName: user.screenName } },
      { type: types.FETCH_USER_SUCCESS, payload: { user } },
    ];
    const store = mockStore({ user: { loading: false, user: null, error: null } });
    return store.dispatch(fetchUser(user.screenName)).then(() => {
      expect(store.getActions()).toEqual(expectedActions);
    });
  });
});
