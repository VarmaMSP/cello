import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { Playlist } from 'types/app'

const playlists: Reducer<{ [playlistId: string]: Playlist }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.RECEIVED_PLAYLISTS:
      return {
        ...state,
        ...action.playlists.reduce<{ [playlistId: string]: Playlist }>(
          (acc, p) => ({ ...acc, [p.id]: p }),
          {},
        ),
      }
    default:
      return state
  }
}

const byUser: Reducer<{ [userId: string]: string[] }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.RECEIVED_PLAYLISTS:
      return {
        ...state,
        [action.userId]: [
          ...new Set([
            ...(state[action.userId] || []),
            ...action.playlists.map((p) => p.id),
          ]),
        ],
      }
    default:
      return state
  }
}

export default combineReducers({
  playlists,
  byUser,
})