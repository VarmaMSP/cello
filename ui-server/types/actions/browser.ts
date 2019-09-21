export const PUSH_PAGE = 'PUSH_PAGE'
export const POP_PAGE = 'POP_PAGE'
export const SET_PAGE_PREVENT_RELOAD = 'SET_PAGE_PREVENT_RELOAD'

export interface PushPageAction {
  type: typeof PUSH_PAGE
  page: { url: string; scrollY: number }
}

export interface PopPageAction {
  type: typeof POP_PAGE
}

export interface SetPagePreventReloadAction {
  type: typeof SET_PAGE_PREVENT_RELOAD
  page: { url: string; scrollY: number }
}

export type BrowserActionTypes =
  | PushPageAction
  | PopPageAction
  | SetPagePreventReloadAction
