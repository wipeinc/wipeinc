export const FETCH_USER_BEGIN = 'FETCH_USER_BEGIN';
export const FETCH_USER_SUCCESS = 'FETCH_USER_SUCCESS';
export const FETCH_USER_FAILURE = 'FETCH_USER_FAILURE';

export const fetchUserBegin = () => ({
  type: FETCH_USER_BEGIN,
});

export const fetchUserSuccess = user => ({
  type: FETCH_USER_SUCCESS,
  payload: { user },
});

export const fetchUserFailure = error => ({
  type: FETCH_USER_FAILURE,
  payload: { error },
});

// Handle HTTP errors since fetch won't.
function checkStatus(response) {
  if (!response.ok) {
    throw Error(response.statusText);
  }
  return response;
}


export default function fetchUser() {
  return (dispatch) => {
    dispatch(fetchUserBegin());
    return fetch('/user')
      .then(checkStatus)
      .then(res => res.json())
      .then((json) => {
        dispatch(fetchUserSuccess(json.USER));
        return json.user;
      })
      .catch(error => dispatch(fetchUserFailure(error)));
  };
}
