import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { Playlist } from 'types/app'
import { addKeyToArr, addToArr } from 'utils/immutable'

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

    case T.PLAYLIST_ADD_EPISODES: {
      let obj = state[action.playlistId]
      return {
        ...state,
        ...(!!obj
          ? {
              [action.playlistId]: {
                ...obj,
                members: addToArr(
                  obj.members,
                  action.episodeIds.map((id, i) => ({
                    episodeId: id,
                    position: obj.members.length + i + 1,
                  })),
                  (x) => x.position,
                ),
              },
            }
          : {}),
      }
    }

    case T.PLAYLIST_REMOVE_EPISODES: {
      let obj = state[action.playlistId]
      return {
        ...state,
        ...(!!obj
          ? {
              [action.playlistId]: {
                ...obj,
                members: obj.members.filter(
                  ({ episodeId }) => !action.episodeIds.includes(episodeId),
                ),
              },
            }
          : {}),
      }
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
      return action.playlists.reduce<{ [userId: string]: string[] }>(
        (acc, p) => ({
          ...acc,
          [p.userId]: addKeyToArr(p.id, acc[p.userId] || []),
        }),
        state,
      )

    default:
      return state
  }
}
 
export default combineReducers({
  byId,
  byUserId,
})
