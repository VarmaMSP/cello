import { AppState } from 'store'

export function getText(state: AppState) {
  return state.ui.searchBar.text
}

export function getIsSearchBarCollapsed(state: AppState) {
  return state.ui.searchBar.collapse
}
