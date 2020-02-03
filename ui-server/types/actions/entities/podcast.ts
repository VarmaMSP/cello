import { Podcast } from 'types/models'

export const PODCAST_ADD = 'podcast/add'

interface AddAction {
  type: typeof PODCAST_ADD
  podcasts: Podcast[]
}

export type PodcastActionTypes =
  | AddAction
