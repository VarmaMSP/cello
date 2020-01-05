import { BrowserActionTypes } from './browser'
import { CurationActionTypes } from './entities/curation'
import { EpisodeActionTypes } from './entities/episode'
import { PlaylistActionTypes } from './entities/playlist'
import { PodcastActionTypes } from './entities/podcast'
import { SearchResultActionTypes } from './entities/search_result'
import { UserActionTypes } from './entities/user'
import { RequestActionTypes } from './request'
import { SessionActionTypes } from './session'
import { AudioPlayerActionTypes } from './ui/audio_player'
import { HistoryFeedActionTypes } from './ui/history_feed'
import { ModalManagerActionTypes } from './ui/modal_manager'
import { PodcastEpisodesListActionTypes } from './ui/podcast_episodes_list'
import { SearchBarActionTypes } from './ui/search_bar'
import { SearchResultsListActionTypes } from './ui/search_results_list'
import { SubscriptionsFeedActionTypes } from './ui/subscriptions_feed'

export type AppActions =
  | UserActionTypes
  | PodcastActionTypes
  | EpisodeActionTypes
  | BrowserActionTypes
  | PlaylistActionTypes
  | RequestActionTypes
  | SessionActionTypes
  | CurationActionTypes
  | AudioPlayerActionTypes
  | HistoryFeedActionTypes
  | SubscriptionsFeedActionTypes
  | ModalManagerActionTypes
  | PodcastEpisodesListActionTypes
  | SearchResultsListActionTypes
  | SearchResultActionTypes
  | SearchBarActionTypes

export * from './browser'
export * from './entities/curation'
export * from './entities/episode'
export * from './entities/feed'
export * from './entities/playlist'
export * from './entities/podcast'
export * from './entities/search_result'
export * from './entities/user'
export * from './request'
export * from './session'
export * from './ui/audio_player'
export * from './ui/history_feed'
export * from './ui/modal_manager'
export * from './ui/podcast_episodes_list'
export * from './ui/search_bar'
export * from './ui/search_results_list'
export * from './ui/subscriptions_feed'
