import { BrowserActionTypes } from './browser'
import { CurationActionTypes } from './curation'
import { PodcastActionTypes } from './podcast'
import { UiActionTypes } from './ui'
import { UserActionTypes } from './user'

export type AppActions =
  | UiActionTypes
  | UserActionTypes
  | PodcastActionTypes
  | CurationActionTypes
  | BrowserActionTypes

export * from './browser'
export * from './curation'
export * from './podcast'
export * from './ui'
export * from './user'
