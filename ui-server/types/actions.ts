import { Podcast, Episode } from './app'

export const GET_PODCAST_REQUEST = 'GET_PODCAST_REQUEST'
export const GET_PODCAST_SUCCESS = 'GET_PODCAST_SUCCESS'
export const GET_PODCAST_FAILURE = 'GET_PODCAST_FAILURE'

export const RECEIVED_PODCAST = 'RECEIVED_PODCAST'
export const RECEIVED_EPISODE = 'RECEIVED_EPISODE'

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

export interface ReceivedPodcastAction {
  type: typeof RECEIVED_PODCAST
  podcast: Podcast
}

export interface ReceivedEpisodeAction {
  type: typeof RECEIVED_EPISODE
  episode: Episode
}

export type PodcastActionTypes =
  | GetPodcastRequestAction
  | GetPodcastSuccessAction
  | GetPodcastFailureAction
  | ReceivedPodcastAction

export type EpisodeActionTypes = ReceivedEpisodeAction

export type AppActions = PodcastActionTypes | EpisodeActionTypes
