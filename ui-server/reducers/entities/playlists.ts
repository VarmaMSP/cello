import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { Playlist } from 'types/app'

const playlists: Reducer<{ [playlistId: string]: Playlist }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.RECEIVED_PLAYLIST:
      return {
        ...state,
        [action.playlist.id]: {
          ...(state[action.playlist.id] || {}),
          ...action.playlist,
        },
      }
    case T.RECEIVED_USER_PLAYLISTS:
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

const playlistsByUser: Reducer<
  {
    [userId: string]: {
      byCreateDateDesc: { [offset: string]: string[] }
      receivedAll: 'create_date_desc'[]
    }
  },
  T.AppActions
> = (state = {}, action) => {
  switch (action.type) {
    case T.RECEIVED_USER_PLAYLISTS:
      switch (action.order) {
        case 'create_date_desc':
          return {
            ...state,
            [action.userId]: {
              ...(state[action.userId] || {}),
              byCreateDateDesc: {
                ...((state[action.userId] || {}).byCreateDateDesc || {}),
                [action.offset.toString()]: action.playlists.map((p) => p.id),
              },
            },
          }
        default:
          return state
      }
    case T.RECEIVED_ALL_USER_PLAYLISTS:
      return {
        ...state,
        [action.userId]: {
          ...(state[action.userId] || {}),
          receivedAll: [
            ...((state[action.userId] || {}).receivedAll || []),
            action.order,
          ],
        },
      }
    default:
      return state
  }
}

export default combineReducers({
  playlists,
  playlistsByUser,
})
