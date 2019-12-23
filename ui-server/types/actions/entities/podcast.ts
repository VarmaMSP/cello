import { Podcast } from 'types/app'

export const PODCAST_ADD = 'podcast/add'

interface AddAction {
  type: typeof PODCAST_ADD
  podcasts: Podcast[]
}

export type PodcastActionTypes =
  | AddAction
