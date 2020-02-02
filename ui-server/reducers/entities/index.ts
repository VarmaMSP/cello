import { combineReducers } from 'redux'
import curations from './curations'
import episodes from './episodes'
import playlists from './playlists'
import podcasts from './podcasts'
import searchResults from './search_results'
import users from './users'

export default combineReducers({
  users,
  podcasts,
  episodes,
  curations,
  playlists,
  searchResults,
})
