import { AppState } from 'store'

export function getActiveModal(state: AppState) {
  return state.ui.modalManager.activeModal
}
