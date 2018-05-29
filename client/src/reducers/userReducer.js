import { Map } from 'immutable';
import * as types from '../actions/actionTypes';
import initialState from './initialState';


export default function userReducer(state = initialState.user, action) {
  switch (action.type) {
    case types.FETCH_USER_BEGIN:
      return Map({
        user: {},
        loading: true,
        error: '',
      });
    case types.FETCH_USER_SUCCESS:
      return Map({
        user: action.payload.user,
        loading: false,
        error: '',
      });
    case types.FETCH_USER_FAILURE:
      return Map({
        user: {},
        loading: false,
        error: action.payload.error,
      });
    default:
      return state;
  }
}
