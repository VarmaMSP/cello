import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { PodcastList } from 'types/app'

const categories: Reducer<{ [id: string]: PodcastList }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.RECEIVED_PODCAST_CATEGORY_LIST:
      return { ...state, [action.category.id]: action.category }
    default:
      return state
  }
}

const subCategories: Reducer<{ [id: string]: string[] }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.RECEIVED_PODCAST_CATEGORY_LIST:
      if (!!!action.category.parentId) {
        return state
      }

      return {
        ...state,
        [action.category.parentId]: [
          ...(state[action.category.parentId] || []),
          action.category.id,
        ],
      }
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
  categories,
  subCategories,
  recommended,
})
