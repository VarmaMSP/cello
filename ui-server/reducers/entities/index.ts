import { combineReducers } from 'redux'
import curations from './curations'
import curationMember from './curation_member'
import episodes from './episodes'
import feed from './feed'
import playlists from './playlists'
import playlistMember from './playlist_member'
import podcasts from './podcasts'
import search from './search'
import users from './users'

export default combineReducers({
  users,
  podcasts,
  episodes,
  curations,
  curationMember,
  playlists,
  playlistMember,
  search,
  feed,
})
