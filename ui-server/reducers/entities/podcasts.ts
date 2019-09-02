import { Podcast } from '../../types/app'
import { RECEIVED_PODCAST } from '../../types/actions'
import { combineReducers, Reducer } from 'redux'

const podcasts: Reducer<{ [PodcastId: string]: Podcast }> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case RECEIVED_PODCAST:
      return { [action.podcast.id]: action.podcast, ...state }
    default:
      return state
  }
}

export default combineReducers({
  podcasts,
})
