import { combineReducers } from 'redux'
import episodes from './episodes'
import feed from './feed'
import playlist from './playlist'
import podcasts from './podcasts'
import search from './search'
import user from './users'

export default combineReducers({
  user,
  podcasts,
  episodes,
  playlist,
  search,
  feed,
})
