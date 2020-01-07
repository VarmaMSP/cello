import { ViewportSize } from 'types/app'

export const WINDOW_VIEWPORT_SIZE = 'window/viewport_size'

export interface ViewportSizeAction {
  type: typeof WINDOW_VIEWPORT_SIZE
  size: ViewportSize
}

export type WindowActionTypes = ViewportSizeAction
