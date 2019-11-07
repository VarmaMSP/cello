import { combineReducers } from 'redux'
import * as T from 'types/actions'
import { defaultRequestReducer } from './utils'

const getUserPlaylists = defaultRequestReducer(
  T.GET_USER_PLAYLISTS_REQUEST,
  T.GET_USER_PLAYLISTS_SUCCESS,
  T.GET_USER_PLAYLISTS_FAILURE,
)

const createPlaylist = defaultRequestReducer(
  T.CREATE_PLAYLIST_REQUEST,
  T.CREATE_PLAYLIST_SUCCESS,
  T.CREATE_PLAYLIST_FAILURE,
)

export default combineReducers({
  getUserPlaylists,
  createPlaylist,
})
