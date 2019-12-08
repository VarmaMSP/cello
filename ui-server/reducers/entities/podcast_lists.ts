import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { PodcastList } from 'types/app'

const podcastsInList: Reducer<{ [id: string]: string[] }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.RECEIVED_PODCASTS_IN_LIST:
      return { ...state, [action.listId]: action.podcasts.map((p) => p.id) }
    default:
      return state
  }
}

const categories: Reducer<{ [id: string]: PodcastList }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.RECEIVED_PODCAST_CATEGORY_LISTS:
      return action.categories.reduce<{ [id: string]: PodcastList }>(
        (acc, c) => ({ ...acc, [c.id]: c }),
        {},
      )
    default:
      return state
  }
}

const subCategories: Reducer<{ [id: string]: string[] }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.RECEIVED_PODCAST_CATEGORY_LISTS:
      return action.categories.reduce<{ [id: string]: string[] }>(
        (acc, c) =>
          !!c.parentId
            ? { ...acc, [c.parentId]: [...(acc[c.parentId] || []), c.id] }
            : acc,
        {},
      )
    default:
      return state
  }
}

const recommended: Reducer<string[], T.AppActions> = (state = [], action) => {
  switch (action.type) {
    case T.RECEIVED_RECOMMENDED_PODCASTS:
      return action.podcasts.map((p) => p.id)
    default:
      return state
  }
}

export default combineReducers({
  podcastsInList,
  categories,
  subCategories,
  recommended,
})
