import { Episode, Podcast } from 'types/app'

export const RECEIVED_PODCAST = 'RECEIVED_PODCAST'
export const RECEIVED_PODCAST_EPISODES = 'RECEIVED_PODCAST_EPISODES'
export const RECEIVED_ALL_PODCAST_EPISODES = 'RECEIVED_ALL_PODCAST_EPISODES'
export const RECEIVED_TRENDING_PODCASTS = 'RECEIVED_TRENDING_PODCASTS'
export const SUBSCRIBED_TO_PODCAST = 'SUBSCRIBED_TO_PODCAST'
export const UNSUBSCRIBED_TO_PODCAST = 'UNSUBSCRIBED_TO_PODCAST'

export interface ReceivedPodcastAction {
  type: typeof RECEIVED_PODCAST
  podcast: Podcast
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

export interface ReceivedTrendingPodcastsAction {
  type: typeof RECEIVED_TRENDING_PODCASTS
  podcasts: Podcast[]
}

export interface SubscribedToPodcastAction {
  type: typeof SUBSCRIBED_TO_PODCAST
  podcastId: string
}

export interface UnsubscribeToPodcastAction {
  type: typeof UNSUBSCRIBED_TO_PODCAST
  podcastId: string
}

export type PodcastActionTypes =
  | ReceivedPodcastAction
  | ReceivedPodcastEpisodesAction
  | ReceivedAllPodcastEpisodesAction
  | ReceivedTrendingPodcastsAction
  | SubscribedToPodcastAction
  | UnsubscribeToPodcastAction
