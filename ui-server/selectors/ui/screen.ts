import { AppState } from '../../store'

export function getScreenWidth(state: AppState) {
  state.ui.screenWidth
}
