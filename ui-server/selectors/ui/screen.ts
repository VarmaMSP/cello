import { AppState } from 'store'

export function getScreenWidth(state: AppState) {
  return state.ui.screenWidth
}
