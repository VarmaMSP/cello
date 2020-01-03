import { combineReducers } from 'redux'
import curations from './curations'
import curationMember from './curation_member'
import episodes from './episodes'
import playlists from './playlists'
import podcasts from './podcasts'
import search from './search'
import searchResults from './search_results'
import users from './users'

export default combineReducers({
  users,
  podcasts,
  episodes,
  curations,
  curationMember,
  playlists,
  search,
  searchResults,
})
