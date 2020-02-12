import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { Podcast } from 'models'

const byId: Reducer<{ [podcastId: string]: Podcast }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.PODCAST_ADD:
      return action.podcasts.reduce<{
        [podcastId: string]: Podcast
      }>(
        (acc, p) => ({ ...acc, [p.id]: Podcast.merge(acc[p.id] || {}, p) }),
        state,
      )

    default:
      return state
  }
}

export default combineReducers({
  byId,
})
