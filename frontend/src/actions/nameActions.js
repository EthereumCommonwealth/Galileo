import { NAME as A } from '../constants/actionTypes'

export function setName(name){
  return {
    type: A.SET_NAME,
    payload: name,
  };
}

export default {
  setName,
}
