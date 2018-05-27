import { Map } from 'immutable';
import * as types from '../actions/actionTypes';
import initialState from './initialState';


export default function userReducer(state = initialState.user, action) {
  switch (action.type) {
    case types.FETCH_USER_BEGIN:
      return Map({
        user: null,
        loading: true,
        error: null,
      });
    case types.FETCH_USER_SUCCESS:
      return Map({
        user: action.payload.user,
        loading: false,
        error: null,
      });
    case types.FETCH_USER_FAILURE:
      return Map({
        user: null,
        loading: false,
        error: action.payload.error,
      });
    default:
      return state;
  }
}
