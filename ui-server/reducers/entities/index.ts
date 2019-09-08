import { combineReducers } from 'redux'
import episodes from './episodes'
import podcasts from './podcasts'
import search from './search'

export default combineReducers({
  podcasts,
  episodes,
  search,
})
