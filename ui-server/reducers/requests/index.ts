import { combineReducers } from 'redux'
import curation from './curation'
import episode from './episode'
import podcast from './podcast'
import search from './search'
import user from './user'

export default combineReducers({
  user,
  podcast,
  episode,
  search,
  curation,
})
