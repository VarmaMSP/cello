import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { Playlist } from 'types/app'
import { addKeyToArr } from 'utils/immutable'

const byId: Reducer<{ [playlistId: string]: Playlist }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.PLAYLIST_ADD:
      return {
        ...state,
        ...action.playlists.reduce<{ [playlistId: string]: Playlist }>(
          (acc, p) => ({ ...acc, [p.id]: { ...(state[p.id] || {}), ...p } }),
          {},
        ),
      }

    default:
      return state
  }
}

const byUserId: Reducer<{ [userId: string]: string[] }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.PLAYLIST_ADD:
      return {
        ...state,
        ...action.playlists.reduce<{ [playlistId: string]: string[] }>(
          (acc, p) => ({
            ...acc,
            [p.id]: addKeyToArr(p.id, state[p.userId] || []),
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
  byUserId,
})
