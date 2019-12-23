import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { PlaylistMember } from 'types/app'
import { addKeysToArr, delKeysFromArr, delKeysFromObj } from 'utils/immutable'

function makeId(playlistId: string, episodeId: string) {
  return `${playlistId}${episodeId}`
}

const byId: Reducer<{ [id: string]: PlaylistMember }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.PLAYLIST_ADD_EPISODES:
      return {
        ...state,
        ...action.episodeIds.reduce<{ [id: string]: PlaylistMember }>(
          (acc, episodeId) => ({
            ...acc,
            [makeIsd(action.playlistId, episodeId)]: {
              id: makeId(action.playlistId, episodeId),
              episodeId: episodeId,
              playlistId: action.playlistId,
            },
          }),
          {},
        ),
      }

    case T.PLAYLIST_REMOVE_EPISODES:
      return delKeysFromObj(
        action.episodeIds.map((episodeId) =>
          makeId(action.playlistId, episodeId),
        ),
        state,
      )

    default:
      return state
  }
}

const byPlaylistId: Reducer<
  { [playlistId: string]: string[] },
  T.AppActions
> = (state = {}, action) => {
  switch (action.type) {
    case T.PLAYLIST_ADD_EPISODES:
      return {
        ...state,
        [action.playlistId]: addKeysToArr(
          action.episodeIds.map((episodeId) =>
            makeId(action.playlistId, episodeId),
          ),
          state[action.playlistId] || [],
        ),
      }

    case T.PLAYLIST_REMOVE_EPISODES:
      return {
        ...state,
        [action.playlistId]: delKeysFromArr(
          action.episodeIds.map((episodeId) =>
            makeId(action.playlistId, episodeId),
          ),
          state[action.playlistId] || [],
        ),
      }

    default:
      return state
  }
}

export default combineReducers({
  byId,
  byPlaylistId,
})
