import { Episode, Podcast } from 'types/app'

export const GET_PODCAST_REQUEST = 'GET_PODCAST_REQUEST'
export const GET_PODCAST_SUCCESS = 'GET_PODCAST_SUCCESS'
export const GET_PODCAST_FAILURE = 'GET_PODCAST_FAILURE'

export const SEARCH_PODCASTS_REQUEST = 'SEARCH_PODCASTS_REQUEST'
export const SEARCH_PODCASTS_SUCCESS = 'SEARCH_PODCASTS_SUCCESS'
export const SEARCH_PODCASTS_FAILURE = 'SEARCH_PODCASTS_FAILURE'

export const RECEIVED_PODCAST = 'RECEIVED_PODCAST'
export const RECEIVED_EPISODES = 'RECEIVED_EPISODES'
export const RECEIVED_SEARCH_PODCASTS = 'RECEIVED_SEARCH_PODCASTS'

export interface GetPodcastRequestAction {
  type: typeof GET_PODCAST_REQUEST
  podcastId: string
}

export interface GetPodcastSuccessAction {
  type: typeof GET_PODCAST_SUCCESS
  podcastId: string
}

export interface GetPodcastFailureAction {
  type: typeof GET_PODCAST_FAILURE
  error: string
}

export interface SearchPodcastRequestAction {
  type: typeof SEARCH_PODCASTS_REQUEST
}

export interface SearchPodcastSuccessAction {
  type: typeof SEARCH_PODCASTS_SUCCESS
}

export interface SearchPodcastFailureAction {
  type: typeof SEARCH_PODCASTS_FAILURE
}

export interface ReceivedPodcastAction {
  type: typeof RECEIVED_PODCAST
  podcast: Podcast
}

export interface ReceivedEpisodesAction {
  type: typeof RECEIVED_EPISODES
  podcastId: string
  episodes: Episode[]
}

export interface ReceivedSearchPodcastsAction {
  type: typeof RECEIVED_SEARCH_PODCASTS
  podcasts: Podcast[]
}

export type PodcastActionTypes =
  | GetPodcastRequestAction
  | GetPodcastSuccessAction
  | GetPodcastFailureAction
  | SearchPodcastRequestAction
  | SearchPodcastSuccessAction
  | SearchPodcastFailureAction
  | ReceivedPodcastAction
  | ReceivedEpisodesAction
  | ReceivedSearchPodcastsAction
