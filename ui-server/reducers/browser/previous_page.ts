import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'

const stack: Reducer<{ urlPath: string; scrollY: number }[], T.AppActions> = (
  state = [],
  action,
) => {
  switch (action.type) {
    case T.PUSH_PREVIOUS_PAGE_STACK:
      return [action.page, ...state]
    case T.POP_PREVIOUS_PAGE_STACK:
      const [, ...pages] = state
      return pages
    default:
      return state
  }
}

const page: Reducer<{ urlPath: string; scrollY: number }, T.AppActions> = (
  state = { urlPath: 'HELLO MORTY', scrollY: 0 },
  action,
) => {
  switch (action.type) {
    case T.SET_PREVIOUS_PAGE:
      return action.page
    default:
      return state
  }
}

export default combineReducers({
  stack,
  page,
})
