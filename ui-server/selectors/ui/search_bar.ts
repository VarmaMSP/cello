import { AppState } from 'store'

export function getText(state: AppState) {
  return state.ui.searchBar.text
}

export function getCursor(state: AppState) {
  return state.ui.searchBar.cursor
}

export function getSuggestions(state: AppState) {
  return state.ui.searchBar.suggestions
}

export function getIsSearchBarCollapsed(state: AppState) {
  return state.ui.searchBar.collapse
}
