import { combineReducers, Reducer } from 'redux'
import { POP_PAGE, PUSH_PAGE, SET_PAGE_PREVENT_RELOAD } from 'types/actions'
import { AppActions } from '../types/actions'

const pages: Reducer<{ url: string; scrollY: number }[], AppActions> = (
  state = [],
  action,
) => {
  switch (action.type) {
    case PUSH_PAGE:
      return [action.page, ...state]
    case POP_PAGE:
      const [, ...pages] = state
      return pages
    default:
      return state
  }
}

const pagePreventReload: Reducer<
  { url: string; scrollY: number },
  AppActions
> = (state = { url: 'HELLO MORTY', scrollY: 0 }, action) => {
  switch (action.type) {
    case SET_PAGE_PREVENT_RELOAD:
      return action.page
    default:
      return state
  }
}

export default combineReducers({
  pages,
  pagePreventReload,
})
