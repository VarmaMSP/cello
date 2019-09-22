import { AppState } from 'store'

export function getCurrentUrlPath(state: AppState) {
  return state.browser.currentUrlPath
}
