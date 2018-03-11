export function setName(name){
  return {
    type: 'SET_NAME',
    payload: name,
  }
}

export default {
  setName,
}
