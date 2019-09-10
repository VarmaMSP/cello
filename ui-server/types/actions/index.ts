import { AppActionTypes } from './app'
import { CurationActionTypes } from './curation'
import { PlayerActionTypes } from './player'
import { PodcastActionTypes } from './podcast'

export type AppActions =
  | PodcastActionTypes
  | CurationActionTypes
  | PlayerActionTypes
  | AppActionTypes

export * from './app'
export * from './curation'
export * from './player'
export * from './podcast'
