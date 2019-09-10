import { Curation, Podcast } from 'types/app'

export const GET_PODCAST_CURATIONS_REQUEST = 'GET_PODCAST_CURATIONS_REQUEST'
export const GET_PODCAST_CURATIONS_SUCCESS = 'GET_PODCAST_CURATIONS_SUCCESS'
export const GET_PODCAST_CURATIONS_FAILURE = 'GET_PODCAST_CURATIONS_FAILURE'
export const RECEIVED_PODCAST_CURATION = 'RECEIVED_PODCAST_CURATION'

export interface GetPodcastCurationsRequestAction {
  type: typeof GET_PODCAST_CURATIONS_REQUEST
}

export interface GetPodcastCurationsSuccessAction {
  type: typeof GET_PODCAST_CURATIONS_SUCCESS
}

export interface GetPodcastCurationsFailureAction {
  type: typeof GET_PODCAST_CURATIONS_FAILURE
  error: string
}

export interface ReceivedPodcastCurationAction {
  type: typeof RECEIVED_PODCAST_CURATION
  curation: Curation
  podcasts: Podcast[]
}

export type CurationActionTypes =
  | GetPodcastCurationsRequestAction
  | GetPodcastCurationsSuccessAction
  | GetPodcastCurationsFailureAction
  | ReceivedPodcastCurationAction
