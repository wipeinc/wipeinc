export const FETCH_PRODUCTS_BEGIN = 'FETCH_PRODUCTS_BEGIN';
export const FETCH_PRODUCTS_SUCCESS = 'FETCH_PRODUCTS_SUCCESS';
export const FETCH_PRODUCTS_FAILURE = 'FETCH_PRODUCTS_FAILURE';

export const fetchUserBegin = () => ({
  type: FETCH_PRODUCTS_BEGIN,
});

export const fetchUserSuccess = user => ({
  type: FETCH_PRODUCTS_SUCCESS,
  payload: { user },
});

export const fetchUserFailure = error => ({
  type: FETCH_PRODUCTS_FAILURE,
  payload: { error },
});

// Handle HTTP errors since fetch won't.
function handleErrors(response) {
  if (!response.ok) {
    throw Error(response.statusText);
  }
  return response;
}

export default function fetchUser() {
  return (dispatch) => {
    dispatch(fetchUserBegin());
    return fetch('/user')
      .then(handleErrors)
      .then(res => res.json())
      .then((json) => {
        dispatch(fetchUserSuccess(json.products));
        return json.user;
      })
      .catch(error => dispatch(fetchUserFailure(error)));
  };
}
