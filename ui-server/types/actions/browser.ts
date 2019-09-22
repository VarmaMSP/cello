import { ViewportSize } from 'types/app'

export const PUSH_PREVIOUS_PAGE_STACK = 'PUSH_PREVIOUS_PAGE_STACK'
export const POP_PREVIOUS_PAGE_STACK = 'POP_PREVIOUS_PAGE_STACK'
export const SET_PREVIOUS_PAGE = 'SET_PREVIOUS_PAGE'
export const SET_VIEWPORT_SIZE = 'SET_VIEWPORT_WIDTH'
export const SET_CURRENT_URL_PATH = 'SET_CURRENT_URL_PATH'

export interface PushPreviousPageStackAction {
  type: typeof PUSH_PREVIOUS_PAGE_STACK
  page: { urlPath: string; scrollY: number }
}

export interface PopPreviousPageStackAction {
  type: typeof POP_PREVIOUS_PAGE_STACK
}

export interface SetPreviousPageAction {
  type: typeof SET_PREVIOUS_PAGE
  page: { urlPath: string; scrollY: number }
}

export interface SetViewportSizeAction {
  type: typeof SET_VIEWPORT_SIZE
  size: ViewportSize
}

export interface SetCurrentUrlPathAction {
  type: typeof SET_CURRENT_URL_PATH
  urlPath: string
}

export type BrowserActionTypes =
  | PushPreviousPageStackAction
  | PopPreviousPageStackAction
  | SetPreviousPageAction
  | SetViewportSizeAction
  | SetCurrentUrlPathAction
