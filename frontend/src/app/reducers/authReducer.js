export default function authenticated(state = {}, { type, payload }) {
  switch (type) {
    case 'SET_LOGGED':
      return payload;
    default:
      return state;
  }
}
