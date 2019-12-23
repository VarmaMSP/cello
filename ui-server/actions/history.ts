import * as client from 'client/history'
import * as T from 'types/actions'
import * as RequestId from 'utils/request_id'
import { requestAction } from './utils'

export function getHistoryPageData() {
  return requestAction(
    () => client.getHistoryPageData(),
    (dispatch, _, { podcasts, episodes }) => {
      dispatch({ type: T.PODCAST_ADD, podcasts })
      dispatch({ type: T.EPISODE_ADD, episodes })
      dispatch({
        type: T.HISTORY_FEED_LOAD_PAGE,
        page: 0,
        episodeIds: episodes.map((x) => x.id),
      })

      if (episodes.length < 10) {
        dispatch({ type: T.HISTORY_FEED_RECEIVED_ALL })
      }
    },
    { requestId: RequestId.getHistoryPageData() },
  )
}

export function getHistoryFeed(offset: number, limit: number) {
  return requestAction(
    () => client.getHistoryFeed(offset, limit),
    (dispatch, _, { podcasts, episodes }) => {
      dispatch({ type: T.PODCAST_ADD, podcasts })
      dispatch({ type: T.EPISODE_ADD, episodes })
      dispatch({
        type: T.HISTORY_FEED_LOAD_PAGE,
        page: Math.floor(offset / 10),
        episodeIds: episodes.map((x) => x.id),
      })

      if (episodes.length < limit) {
        dispatch({ type: T.HISTORY_FEED_RECEIVED_ALL })
      }
    },
    { requestId: RequestId.getHistoryFeed() },
  )
}
