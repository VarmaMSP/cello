import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { PodcastEpisodeListOrder } from 'types/ui'

const list: Reducer<
  {
    [podcastId: string]: {
      [order in PodcastEpisodeListOrder]: {
        [page: number]: string[]
      }
    }
  },
  T.AppActions
> = (state = {}, action) => {
  switch (action.type) {
    case T.PODCAST_EPISODES_LIST_LOAD_PAGE:
      return {
        ...state,
        [action.podcastId]: {
          ...(state[action.podcastId] || {}),
          [action.order]: {
            ...((state[action.podcastId] || {})[action.order] || {}),
            [action.page]: action.episodeIds,
          },
        },
      }

    default:
      return state
  }
}

const receivedAll: Reducer<
  {
    [podcastId: string]: PodcastEpisodeListOrder[]
  },
  T.AppActions
> = (state = {}, action) => {
  switch (action.type) {
    case T.PODCAST_EPISODES_LIST_RECEIVED_ALL:
      return {
        ...state,
        [action.podcastId]: [...(state[action.podcastId] || []), action.order],
      }

    default:
      return state
  }
}

export default combineReducers({
  list,
  receivedAll,
})
