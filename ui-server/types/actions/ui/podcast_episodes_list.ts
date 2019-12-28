import { PodcastEpisodeListOrder } from 'types/ui'

export const PODCAST_EPISODES_LIST_LOAD_PAGE = 'podcast_episodes_list/load_page'
export const PODCAST_EPISODES_LIST_RECEIVED_ALL =
  'podcast_episodes_list/received_all'

interface LoadPageAction {
  type: typeof PODCAST_EPISODES_LIST_LOAD_PAGE
  podcastId: string
  episodeIds: string[]
  page: number
  order: PodcastEpisodeListOrder
}

interface ReceivedAllAction {
  type: typeof PODCAST_EPISODES_LIST_RECEIVED_ALL
  podcastId: string
  order: PodcastEpisodeListOrder
}

export type PodcastEpisodesListActionTypes = LoadPageAction | ReceivedAllAction
