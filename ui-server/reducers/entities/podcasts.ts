import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { Podcast } from 'types/app'

const podcasts: Reducer<{ [PodcastId: string]: Podcast }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.RECEIVED_PODCAST:
      return { ...state, [action.podcast.id]: action.podcast }
    case T.RECEIVED_SEARCH_PODCASTS:
    case T.RECEIVED_PODCAST_CURATION:
    case T.RECEIVED_TRENDING_PODCASTS:
      return {
        ...state,
        ...action.podcasts.reduce<{
          [id: string]: Podcast
        }>((acc, p) => ({ ...acc, [p.id]: p }), {}),
      }
    case T.RECEIVED_SIGNED_IN_USER:
      return {
        ...action.subscriptions.reduce<{
          [id: string]: Podcast
        }>((acc, p) => ({ ...acc, [p.id]: p }), {}),
        ...state,
      }
    default:
      return state
  }
}

const podcastsTrending: Reducer<string[], T.AppActions> = (
  state = [],
  action,
) => {
  switch (action.type) {
    case T.RECEIVED_TRENDING_PODCASTS:
      return action.podcasts.map((p) => p.id)
    default:
      return state
  }
}

const podcastsInCuration: Reducer<
  { [curationId: string]: string[] },
  T.AppActions
> = (state = {}, action) => {
  switch (action.type) {
    case T.RECEIVED_PODCAST_CURATION:
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
  T.AppActions
> = (state = {}, action) => {
  switch (action.type) {
    case T.RECEIVED_SIGNED_IN_USER:
      return {
        ...state,
        [action.user.id]: action.subscriptions.map((p) => p.id),
      }
    case T.SUBSCRIBED_TO_PODCAST:
      return {
        ...state,
        [action.userId]: [action.podcastId],
      }
    case T.UNSUBSCRIBED_TO_PODCAST:
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
  podcastsTrending,
  podcastsInCuration,
  podcastsSubscribedByUser,
})
