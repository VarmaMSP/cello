import { combineReducers } from 'redux'
import curations from './curations'
import episodes from './episodes'
import podcasts from './podcasts'
import search from './search'

export default combineReducers({
  podcasts,
  episodes,
  curations,
  search,
})
