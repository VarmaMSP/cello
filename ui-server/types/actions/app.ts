import { ScreenWidth } from 'types/app'

export const SET_SCREEN_WIDTH = 'SET_SCREEN_WIDTH'

export interface SetScreenWidthAction {
  type: typeof SET_SCREEN_WIDTH
  width: ScreenWidth
}

export type AppActionTypes = SetScreenWidthAction
