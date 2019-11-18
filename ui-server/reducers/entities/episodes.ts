import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { Episode, EpisodePlayback } from 'types/app'

const episodes: Reducer<{ [episodeId: string]: Episode }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.RECEIVED_EPISODES:
    case T.RECEIVED_PODCAST_EPISODES:
    case T.RECEIVED_USER_FEED:
    case T.RECEIVED_USER_FEED_PUBLISHED_BEFORE:
      return {
        ...state,
        ...action.episodes.reduce<{ [id: string]: Episode }>(
          (acc, e) => ({ ...acc, [e.id]: e }),
          {},
        ),
      }
    default:
      return state
  }
}

const episodesInPodcast: Reducer<
  {
    [podcastId: string]: {
      byPubDateDesc: { [offset: string]: string[] }
      byPubDateAsc: { [offset: string]: string[] }
    }
  },
  T.AppActions
> = (state = {}, action) => {
  switch (action.type) {
    case T.RECEIVED_PODCAST_EPISODES:
      switch (action.order) {
        case 'PUB_DATE_DESC':
          return {
            ...state,
            [action.podcastId]: {
              ...(state[action.podcastId] || {}),
              byPubDateDesc: {
                ...((state[action.podcastId] || {}).byPubDateDesc || {}),
                [action.offset.toString()]: action.episodes.map((e) => e.id),
              },
            },
          }
        case 'PUB_DATE_DESC':
          return {
            ...state,
            [action.podcastId]: {
              ...(state[action.podcastId] || {}),
              byPubDateAsc: {
                ...((state[action.podcastId] || {}).byPubDateAsc || {}),
                [action.offset.toString()]: action.episodes.map((e) => e.id),
              },
            },
          }
      }
    default:
      return state
  }
}

const currentUserFeedPublishedBefore: Reducer<
  { [publishedBefore: string]: string[] },
  T.AppActions
> = (state = {}, action) => {
  switch (action.type) {
    case T.RECEIVED_USER_FEED_PUBLISHED_BEFORE:
      return {
        ...state,
        [action.publishedBefore]: action.episodes.map((e) => e.id),
      }
    default:
      return state
  }
}

const currentUserHistory: Reducer<string[], T.AppActions> = (
  state = [],
  action,
) => {
  switch (action.type) {
    case T.RECEIVED_HISTORY_PLAYBACKS:
      return action.playbacks.map((e) => e.id)
    default:
      return state
  }
}

const currentUserPlayback: Reducer<
  { [episodeId: string]: EpisodePlayback },
  T.AppActions
> = (state = {}, action) => {
  switch (action.type) {
    case T.RECEIVED_HISTORY_PLAYBACKS:
    case T.RECEIVED_EPISODE_PLAYBACKS:
      return {
        ...state,
        ...action.playbacks.reduce<{
          [episodeId: string]: EpisodePlayback
        }>((acc, playback) => ({ ...acc, [playback.episodeId]: playback }), {}),
      }
    default:
      return state
  }
}

export default combineReducers({
  episodes,
  episodesInPodcast,
  currentUserFeed: combineReducers({
    publishedBefore: currentUserFeedPublishedBefore,
  }),
  currentUserHistory,
  currentUserPlayback,
})
