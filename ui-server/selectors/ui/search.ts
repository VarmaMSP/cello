import { AppState } from 'store'

export function getSearchBarText(state: AppState) {
  return state.ui.searchText
}
