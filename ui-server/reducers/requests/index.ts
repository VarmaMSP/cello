import { combineReducers } from 'redux'
import curation from './curation'
import podcast from './podcast'
import search from './search'
import user from './user'

export default combineReducers({
  user,
  podcast,
  search,
  curation,
})
