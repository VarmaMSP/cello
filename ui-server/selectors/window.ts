import { AppState } from 'store'

export function getViewportSize(state: AppState) {
  return state.window.viewportSize
}
