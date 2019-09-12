import { CurationActionTypes } from './curation'
import { PodcastActionTypes } from './podcast'
import { UiActionTypes } from './ui'

export type AppActions =
  | UiActionTypes
  | PodcastActionTypes
  | CurationActionTypes

export * from './curation'
export * from './podcast'
export * from './ui'
