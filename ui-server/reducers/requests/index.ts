import { combineReducers } from 'redux'
import curation from './curation'
import podcast from './podcast'
import search from './search'

export default combineReducers({
  podcast,
  search,
  curation,
})
