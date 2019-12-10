import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { Podcast } from 'types/app'

const podcasts: Reducer<{ [PodcastId: string]: Podcast }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.RECEIVED_PODCAST:
      return {
        ...state,
        [action.podcast.id]: {
          ...(state[action.podcast.id] || {}),
          ...action.podcast,
        },
      }
    case T.RECEIVED_PODCASTS:
    case T.RECEIVED_CHART_PODCASTS:
    case T.RECEIVED_SEARCH_PODCASTS:
      return {
        ...state,
        ...action.podcasts.reduce<{
          [id: string]: Podcast
        }>(
          (acc, p) => ({ ...acc, [p.id]: { ...(state[p.id] || {}), ...p } }),
          {},
        ),
      }
    case T.RECEIVED_SIGNED_IN_USER:
      return {
        ...state,
        ...action.subscriptions.reduce<{
          [id: string]: Podcast
        }>(
          (acc, p) => ({ ...acc, [p.id]: { ...(state[p.id] || {}), ...p } }),
          {},
        ),
      }
    default:
      return state
  }
}

const subscriptions: Reducer<string[], T.AppActions> = (state = [], action) => {
  switch (action.type) {
    case T.RECEIVED_SIGNED_IN_USER:
      return action.subscriptions.map((p) => p.id)
    case T.SUBSCRIBED_TO_PODCAST:
      return [...new Set([action.podcastId, ...state])]
    case T.UNSUBSCRIBED_TO_PODCAST:
      return state.filter((id) => id !== action.podcastId)
    default:
      return state
  }
}

export default combineReducers({
  podcasts,
  subscriptions,
})
