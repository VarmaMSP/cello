import * as client from 'client/history'
import * as T from 'types/actions'
import * as RequestId from 'utils/request_id'
import { requestAction } from './utils'

export function getHistoryPageData() {
  return requestAction(
    () => client.getHistoryFeed(0, 15),
    (dispatch, _, { podcasts, episodes }) => {
      dispatch({
        type: T.RECEIVED_PODCASTS,
        podcasts,
      })
      dispatch({
        type: T.RECEIVED_HISTORY_FEED,
        offset: 0,
        episodes,
      })

      if (episodes.length < 15) {
        dispatch({ type: T.RECEIVED_ALL_HISTORY_FEED })
      }
    },
    { requestId: RequestId.getHistoryPageData() },
  )
}

export function getHistoryFeed(offset: number, limit: number) {
  return requestAction(
    () => client.getHistoryFeed(offset, limit),
    (dispatch, _, { podcasts, episodes }) => {
      dispatch({
        type: T.RECEIVED_PODCASTS,
        podcasts,
      })
      dispatch({
        type: T.RECEIVED_HISTORY_FEED,
        offset,
        episodes,
      })

      if (episodes.length < limit) {
        dispatch({ type: T.RECEIVED_ALL_HISTORY_FEED })
      }
    },
    { requestId: RequestId.getHistoryFeed() },
  )
}
