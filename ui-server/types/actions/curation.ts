import { Curation, Podcast } from 'types/app'

export const RECEIVED_PODCAST_CURATION = 'RECEIVED_PODCAST_CURATION'

export interface ReceivedPodcastCurationAction {
  type: typeof RECEIVED_PODCAST_CURATION
  curation: Curation
  podcasts: Podcast[]
}

export type CurationActionTypes = ReceivedPodcastCurationAction
