import * as types from './actionTypes';

export const fetchUserBegin = screenName => ({
  type: types.FETCH_USER_BEGIN,
  payload: { screenName },
});

export const fetchUserSuccess = user => ({
  type: types.FETCH_USER_SUCCESS,
  payload: { user },
});

export const fetchUserFailure = error => ({
  type: types.FETCH_USER_FAILURE,
  payload: { error },
});

// Handle HTTP errors since fetch won't.
function checkStatus(response) {
  if (!response.ok) {
    throw Error(response.statusText);
  }
  return response;
}

export default function fetchUser(screenName) {
  console.log(`fetching ${screenName}`);
  return (dispatch) => {
    dispatch(fetchUserBegin(screenName));
    return fetch(`${__API_BASE_URL__}/profile/${screenName}`)
      .then(checkStatus)
      .then(res => res.json())
      .then((data) => {
        dispatch(fetchUserSuccess(data));
      })
      .catch(error => dispatch(fetchUserFailure(error)));
  };
}
