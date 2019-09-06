import { ScreenWidth } from 'types/app'

export const SET_SCREEN_WIDTH = 'SET_SCREEN_WIDTH'
export const SEARCH_BAR_TEXT_CHANGE = 'SEARCH_BAR_TEXT_CHANGE'

export interface SetScreenWidthAction {
  type: typeof SET_SCREEN_WIDTH
  width: ScreenWidth
}

export interface SearchBarTextChangeAction {
  type: typeof SEARCH_BAR_TEXT_CHANGE
  text: string
}

export type AppActionTypes = SetScreenWidthAction | SearchBarTextChangeAction
