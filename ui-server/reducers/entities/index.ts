import { combineReducers } from 'redux'
import curations from './curations'
import episodes from './episodes'
import podcasts from './podcasts'
import search from './search'
import user from './users'

export default combineReducers({
  user,
  podcasts,
  episodes,
  curations,
  search,
})
