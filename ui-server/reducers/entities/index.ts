import { combineReducers } from 'redux'
import podcasts from './podcasts'
import episodes from './episodes'
import search from './search'

export default combineReducers({
  podcasts,
  episodes,
  search,
})
