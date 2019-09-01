import { Podcast } from 'types/app'
import { AppActions, RECEIVED_PODCAST } from 'types/actions'
import { combineReducers } from 'redux'

type PodcastsState = { [key: string]: Podcast }

const podcasts = (
  state: PodcastsState = {},
  action: AppActions,
): PodcastsState => {
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
