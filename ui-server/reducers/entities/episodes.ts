import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { Episode } from 'types/app'
import { addKeyToArr } from 'utils/immutable'

const byId: Reducer<{ [episodeId: string]: Episode }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.EPISODE_ADD:
      return {
        ...state,
        ...action.episodes.reduce<{ [episodeId: string]: Episode }>(
          (acc, e) => ({ ...acc, [e.id]: { ...(state[e.id] || {}), ...e } }),
          {},
        ),
      }

    case T.EPISODE_JOIN_PLAYBACK:
      return {
        ...state,
        ...action.playbacks.reduce<{ [episodeId: string]: Episode }>(
          (acc, p) => ({
            ...acc,
            [p.episodeId]: { ...(state[p.episodeId] || {}), ...p },
          }),
          {},
        ),
      }

    default:
      return state
  }
}

const byPodcastId: Reducer<{ [podcastId: string]: string[] }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.EPISODE_ADD:
      return {
        ...state,
        ...action.episodes.reduce<{ [podcastId: string]: string[] }>(
          (acc, e) => ({
            ...acc,
            [e.podcastId]: addKeyToArr(e.id, state[e.podcastId] || []),
          }),
          {},
        ),
      }

    default:
      return state
  }
}

export default combineReducers({
  byId,
  byPodcastId,
})
