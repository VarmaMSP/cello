import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { Episode } from 'types/app'

const episodes: Reducer<{ [episodeId: string]: Episode }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.RECEIVED_EPISODES:
    case T.RECEIVED_USER_FEED:
      const episodes = action.episodes.reduce<{ [id: string]: Episode }>(
        (acc, e) => ({ ...acc, [e.id]: e }),
        {},
      )
      return { ...state, ...episodes }
    default:
      return state
  }
}

const episodesInPodcast: Reducer<
  { [podcastId: string]: string[] },
  T.AppActions
> = (state = {}, action) => {
  switch (action.type) {
    case T.RECEIVED_EPISODES:
      return {
        ...state,
        [action.podcastId]: action.episodes.map((e) => e.id),
      }
    default:
      return state
  }
}

export default combineReducers({
  episodes,
  episodesInPodcast,
})
