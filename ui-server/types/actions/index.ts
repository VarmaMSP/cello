import { BrowserActionTypes } from './browser'
import { EpisodeActionTypes } from './entities/episode'
import { FeedActionTypes } from './entities/feed'
import { PlaylistActionTypes } from './entities/playlist'
import { PodcastActionTypes } from './entities/podcast'
import { PodcastListsActionTypes } from './entities/podcast_lists'
import { SearchActionTypes } from './entities/search'
import { UserActionTypes } from './entities/user'
import { RequestActionTypes } from './request'
import { UiActionTypes } from './ui'

export type AppActions =
  | UiActionTypes
  | UserActionTypes
  | PodcastActionTypes
  | EpisodeActionTypes
  | BrowserActionTypes
  | PlaylistActionTypes
  | RequestActionTypes
  | SearchActionTypes
  | FeedActionTypes
  | PodcastListsActionTypes

export * from './browser'
export * from './entities/episode'
export * from './entities/feed'
export * from './entities/playlist'
export * from './entities/podcast'
export * from './entities/podcast_lists'
export * from './entities/search'
export * from './entities/user'
export * from './request'
export * from './ui'
