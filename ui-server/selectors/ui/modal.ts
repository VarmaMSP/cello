import { AppState } from 'store'

export function getModalToShow(state: AppState) {
  return state.ui.showModal
}
