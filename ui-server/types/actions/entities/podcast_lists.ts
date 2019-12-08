import { Podcast, PodcastList } from 'types/app'

export const RECEIVED_RECOMMENDED_PODCASTS = 'RECEIVED_RECOMMENDED_PODCASTS'
export const RECEIVED_PODCAST_CATEGORY_LISTS = 'RECEIVED_PODCAST_CATEGORY_LISTS'

export interface ReceivedRecommendedPodcastsAction {
  type: typeof RECEIVED_RECOMMENDED_PODCASTS
  podcasts: Podcast[]
}

export interface ReceivedPodcastCategoryListsAction {
  type: typeof RECEIVED_PODCAST_CATEGORY_LISTS
  categories: PodcastList[]
}

export type PodcastListsActionTypes =
  | ReceivedRecommendedPodcastsAction
  | ReceivedPodcastCategoryListsAction
