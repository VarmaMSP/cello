import { combineReducers } from 'redux'
import curation from './curation'
import episode from './episode'
import playlist from './playlist'
import podcast from './podcast'
import search from './search'
import user from './user'

export default combineReducers({
  user,
  podcast,
  episode,
  search,
  curation,
  playlist,
})
