import { Podcast, PodcastList } from 'types/app'

export const RECEIVED_RECOMMENDED_PODCASTS = 'RECEIVED_RECOMMENDED_PODCASTS'
export const RECEIVED_PODCAST_CATEGORY_LIST = 'RECEIVED_PODCAST_CATEGORY_LIST'

export interface ReceivedRecommendedPodcastsAction {
  type: typeof RECEIVED_RECOMMENDED_PODCASTS
  podcasts: Podcast[]
}

export interface ReceivedPodcastCategoryListAction {
  type: typeof RECEIVED_PODCAST_CATEGORY_LIST
  category: PodcastList
}

export type PodcastListsActionTypes =
  | ReceivedRecommendedPodcastsAction
  | ReceivedPodcastCategoryListAction
