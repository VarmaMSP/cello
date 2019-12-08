import { combineReducers } from 'redux'
import charts from './charts'
import episodes from './episodes'
import feed from './feed'
import playlists from './playlists'
import podcasts from './podcasts'
import search from './search'
import user from './users'

export default combineReducers({
  user,
  podcasts,
  episodes,
  playlists,
  search,
  feed,
  charts,
})
