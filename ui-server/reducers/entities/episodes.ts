import { Episode } from '../../types/app'
import { AppActions, RECEIVED_EPISODES } from '../../types/actions'
import { combineReducers, Reducer } from 'redux'

export const episodes: Reducer<{ [episodeId: string]: Episode }, AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case RECEIVED_EPISODES:
      const episodes = action.episodes.reduce<{ [id: string]: Episode }>(
        (acc, e) => ({ ...acc, [e.id]: e }),
        {},
      )
      return { ...state, ...episodes }
    default:
      return state
  }
}

export const episodesInPodcast: Reducer<
  { [podcastId: string]: string[] },
  AppActions
> = (state = {}, action) => {
  switch (action.type) {
    case RECEIVED_EPISODES:
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
