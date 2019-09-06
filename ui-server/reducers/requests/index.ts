import { combineReducers } from 'redux'
import podcast from './podcast'
import search from './search'

export default combineReducers({
  podcast,
  search,
})
