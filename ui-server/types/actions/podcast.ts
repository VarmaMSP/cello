import { Episode, Podcast } from 'types/app'

export const RECEIVED_PODCAST = 'RECEIVED_PODCAST'
export const RECEIVED_EPISODES = 'RECEIVED_EPISODES'
export const RECEIVED_PODCAST_EPISODES = 'RECEIVED_PODCAST_EPISODES'
export const RECEIVED_ALL_PODCAST_EPISODES = 'RECEIVED_ALL_PODCAST_EPISODES'
export const RECEIVED_TRENDING_PODCASTS = 'RECEIVED_TRENDING_PODCASTS'
export const RECEIVED_SEARCH_PODCASTS = 'RECEIVED_SEARCH_PODCASTS'

export interface ReceivedPodcastAction {
  type: typeof RECEIVED_PODCAST
  podcast: Podcast
}

export interface ReceivedTrendingPodcastsAction {
  type: typeof RECEIVED_TRENDING_PODCASTS
  podcasts: Podcast[]
}

export interface ReceivedEpisodesAction {
  type: typeof RECEIVED_EPISODES
  podcastId: string
  episodes: Episode[]
}

export interface ReceivedPodcastEpisodesAction {
  type: typeof RECEIVED_PODCAST_EPISODES
  podcastId: string
  order: 'pub_date_desc' | 'pub_date_asc'
  offset: number
  episodes: Episode[]
}

export interface ReceivedAllPodcastEpisodesAction {
  type: typeof RECEIVED_ALL_PODCAST_EPISODES
  podcastId: string
  order: 'pub_date_desc' | 'pub_date_asc'
}

export interface ReceivedSearchPodcastsAction {
  type: typeof RECEIVED_SEARCH_PODCASTS
  query: string
  podcasts: Podcast[]
}

export type PodcastActionTypes =
  | ReceivedPodcastAction
  | ReceivedTrendingPodcastsAction
  | ReceivedEpisodesAction
  | ReceivedPodcastEpisodesAction
  | ReceivedAllPodcastEpisodesAction
  | ReceivedSearchPodcastsAction
