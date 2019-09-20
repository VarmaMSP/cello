import { combineReducers, Reducer } from 'redux'
import {
  AppActions,
  RECEIVED_PODCAST,
  RECEIVED_PODCAST_CURATION,
  RECEIVED_SEARCH_PODCASTS,
  SUBSCRIBED_TO_PODCAST,
  UNSUBSCRIBED_TO_PODCAST,
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
    case RECEIVED_PODCAST_CURATION:
      const podcasts = action.podcasts.reduce<{ [id: string]: Podcast }>(
        (acc, p) => ({ ...acc, [p.id]: p }),
        {},
      )
      return { ...state, ...podcasts }
    default:
      return state
  }
}

const podcastsInCuration: Reducer<
  { [curationId: string]: string[] },
  AppActions
> = (state = {}, action) => {
  switch (action.type) {
    case RECEIVED_PODCAST_CURATION:
      return {
        ...state,
        [action.curation.id]: action.podcasts.map((p) => p.id),
      }
    default:
      return state
  }
}

const podcastsSubscribedByUser: Reducer<
  { [userId: string]: string[] },
  AppActions
> = (state = {}, action) => {
  switch (action.type) {
    case SUBSCRIBED_TO_PODCAST:
      return {
        ...state,
        [action.userId]: [action.podcastId],
      }
    case UNSUBSCRIBED_TO_PODCAST:
      return {
        ...state,
        [action.userId]: (state[action.userId] || []).filter(
          (id) => id !== action.podcastId,
        ),
      }
    default:
      return state
  }
}

export default combineReducers({
  podcasts,
  podcastsInCuration,
  podcastsSubscribedByUser,
})
