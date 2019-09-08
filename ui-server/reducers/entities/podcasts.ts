import { combineReducers, Reducer } from 'redux'
import {
  AppActions,
  RECEIVED_PODCAST,
  RECEIVED_SEARCH_PODCASTS,
} from 'types/actions'
import { Podcast } from 'types/app'

const podcasts: Reducer<{ [PodcastId: string]: Podcast }, AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case RECEIVED_PODCAST:
      return { ...state, [action.podcast.id]: action.podcast }
    case RECEIVED_SEARCH_PODCASTS:
      const podcasts = action.podcasts.reduce<{ [id: string]: Podcast }>(
        (acc, p) => ({ ...acc, [p.id]: p }),
        {},
      )
      return { ...state, ...podcasts }
    default:
      return state
  }
}

export default combineReducers({
  podcasts,
})
