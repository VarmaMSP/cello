import { BrowserActionTypes } from './browser'
import { EpisodeActionTypes } from './episode'
import { FeedActionTypes } from './feed'
import { PlaylistActionTypes } from './playlist'
import { PodcastActionTypes } from './podcast'
import { RequestActionTypes } from './request'
import { UiActionTypes } from './ui'
import { UserActionTypes } from './user'

export type AppActions =
  | UiActionTypes
  | UserActionTypes
  | PodcastActionTypes
  | EpisodeActionTypes
  | BrowserActionTypes
  | PlaylistActionTypes
  | RequestActionTypes
  | FeedActionTypes

export * from './browser'
export * from './episode'
export * from './feed'
export * from './playlist'
export * from './podcast'
export * from './request'
export * from './ui'
export * from './user'
