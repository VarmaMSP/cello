import { AppState } from 'store'

export function getText(state: AppState) {
  return state.ui.searchBar.text
}
