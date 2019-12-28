import { BrowserActionTypes } from './browser'
import { CurationActionTypes } from './entities/curation'
import { EpisodeActionTypes } from './entities/episode'
import { PlaylistActionTypes } from './entities/playlist'
import { PodcastActionTypes } from './entities/podcast'
import { SearchActionTypes } from './entities/search'
import { UserActionTypes } from './entities/user'
import { RequestActionTypes } from './request'
import { SessionActionTypes } from './session'
import { UiActionTypes } from './ui'
import { AudioPlayerActionTypes } from './ui/audio_player'
import { HistoryFeedActionTypes } from './ui/history_feed'
import { ModalManagerActionTypes } from './ui/modal_manager'
import { PodcastEpisodesListActionTypes } from './ui/podcast_episodes_list'
import { SubscriptionsFeedActionTypes } from './ui/subscriptions_feed'

export type AppActions =
  | UiActionTypes
  | UserActionTypes
  | PodcastActionTypes
  | EpisodeActionTypes
  | BrowserActionTypes
  | PlaylistActionTypes
  | RequestActionTypes
  | SearchActionTypes
  | SessionActionTypes
  | CurationActionTypes
  | AudioPlayerActionTypes
  | HistoryFeedActionTypes
  | SubscriptionsFeedActionTypes
  | ModalManagerActionTypes
  | PodcastEpisodesListActionTypes

export * from './browser'
export * from './entities/curation'
export * from './entities/episode'
export * from './entities/feed'
export * from './entities/playlist'
export * from './entities/podcast'
export * from './entities/search'
export * from './entities/user'
export * from './request'
export * from './session'
export * from './ui'
export * from './ui/audio_player'
export * from './ui/history_feed'
export * from './ui/modal_manager'
export * from './ui/podcast_episodes_list'
export * from './ui/subscriptions_feed'
