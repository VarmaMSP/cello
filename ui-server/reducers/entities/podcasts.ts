import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { Podcast } from 'types/models'

const byId: Reducer<{ [podcastId: string]: Podcast }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.PODCAST_ADD:
      return {
        ...state,
        ...action.podcasts.reduce<{
          [podcastId: string]: Podcast
        }>(
          (acc, p) => ({ ...acc, [p.id]: { ...(state[p.id] || {}), ...p } }),
          {},
        ),
      }

    default:
      return state
  }
}

export default combineReducers({
  byId,
})
