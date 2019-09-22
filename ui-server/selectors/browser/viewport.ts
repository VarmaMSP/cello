import { AppState } from 'store'

export function getViewportSize(state: AppState) {
  return state.browser.viewportSize
}
