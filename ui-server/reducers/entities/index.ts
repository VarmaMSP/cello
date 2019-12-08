import { combineReducers } from 'redux'
import episodes from './episodes'
import feed from './feed'
import playlists from './playlists'
import podcasts from './podcasts'
import podcastLists from './podcast_lists'
import search from './search'
import user from './users'
export default combineReducers({
  user,
  podcasts,
  episodes,
  playlists,
  search,
  feed,
  podcastLists,
})
