import { Category } from 'models/category'

export const CATEGORY_ADD = 'category/add'
export const CATEGORY_ADD_PODCASTS = 'category/add_podcasts'

interface AddAction {
  type: typeof CATEGORY_ADD
  categories: Category[]
}

interface AddPodcastsAction {
  type: typeof CATEGORY_ADD_PODCASTS
  categoryId: string
  podcastIds: string[]
}

export type CategoryActionTypes = AddAction | AddPodcastsAction
