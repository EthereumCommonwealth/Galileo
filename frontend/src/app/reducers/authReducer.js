import initialState from '../initialState';

export default function authenticated(state = initialState, { type, payload }) {
  switch (type) {
    case 'SET_LOGGED':
      return payload;
    default:
      return state;
  }
}
