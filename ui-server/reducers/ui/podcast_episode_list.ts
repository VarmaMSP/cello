import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { PodcastEpisodeListSortOrder } from 'types/ui'
import { PODCAST_EPISODES_LIST_LOAD_PAGE } from 'types/actions'

const listByPubDateDesc: Reducer<
  { 
    [podcastId: string]
  }

const list: Reducer<
  {
    [podcastId: string]: {
      [order in PodcastEpisodeListSortOrder]: {
        [page: number]: string[]
      }
    }
  },
  T.AppActions
> = (state = {}, action) => {
  switch (action.type) {
    case PODCAST_EPISODES_LIST_LOAD_PAGE:
      return {
        ...state,
        [action.podcastId]: {
          ...(state[action.podcastId] || {}),
          [action.order]: {
            ...((state[action.podcastId] || {})[action.order] || {}),
            [action.page]: action
          }
        }
      }
    default:
      return state
  }
}

const receivedAll: Reducer<PodcastEpisodeListSortOrder[], T.AppActions> = (
  state = [],
  action,
) => {
  switch (action.type) {
    default:
      return state
  }
}

export default combineReducers({
  list,
  receivedAll,
})
