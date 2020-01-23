import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'

const text: Reducer<string, T.AppActions> = (state = '', action) => {
  switch (action.type) {
    case T.SEARCH_BAR_UPDATE_TEXT:
      return action.text

    default:
      return state
  }
}

const collapse: Reducer< boolean, T.AppActions> = (state = true, action) => {
  switch (action.type) {
    case T.SEARCH_BAR_EXPAND:
      return false
    
    case T.SEARCH_BAR_COLLAPSE:
      return true

    default:
      return state
  }
}

export default combineReducers({
  text,
  collapse,
})
