import { Podcast } from 'types/app'

export const RECEIVED_SEARCH_PODCASTS = 'RECEIVED_SEARCH_PODCASTS'

export interface ReceivedSearchPodcastsAction {
  type: typeof RECEIVED_SEARCH_PODCASTS
  query: string
  podcasts: Podcast[]
}

export type SearchActionTypes = ReceivedSearchPodcastsAction
