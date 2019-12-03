import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { Episode } from 'types/app'

const episodes: Reducer<{ [episodeId: string]: Episode }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.RECEIVED_EPISODES:
    case T.RECEIVED_PODCAST_EPISODES:
    case T.RECEIVED_SUBSCRIPTION_FEED:
    case T.RECEIVED_HISTORY_FEED:
      return {
        ...state,
        ...action.episodes.reduce<{ [episodeId: string]: Episode }>(
          (acc, e) => ({ ...acc, [e.id]: e }),
          {},
        ),
      }
    case T.RECEIVED_PLAYBACKS:
      return {
        ...state,
        ...action.playbacks.reduce<{ [episodeId: string]: Episode }>(
          (acc, p) => ({
            ...acc,
            [p.episodeId]: { ...state[p.episodeId], ...p },
          }),
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
      receivedAll: ('pub_date_desc' | 'pub_date_asc')[]
    }
  },
  T.AppActions
> = (state = {}, action) => {
  switch (action.type) {
    case T.RECEIVED_PODCAST_EPISODES:
      switch (action.order) {
        case 'pub_date_desc':
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
        case 'pub_date_asc':
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
    case T.RECEIVED_ALL_PODCAST_EPISODES:
      return {
        ...state,
        [action.podcastId]: {
          ...(state[action.podcastId] || {}),
          receivedAll: [
            ...((state[action.podcastId] || {}).receivedAll || []),
            action.order,
          ],
        },
      }
    default:
      return state
  }
}

export default combineReducers({
  episodes,
  episodesInPodcast,
})
