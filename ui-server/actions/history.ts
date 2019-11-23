import client from 'client'
import * as T from 'types/actions'
import * as RequestId from 'utils/request_id'
import { requestAction } from './utils'

export function getHistoryFeed(offset: number, limit: number) {
  return requestAction(
    () => client.getHistoryFeed(offset, limit),
    (dispatch, _, { episodes, playbacks }) => {
      dispatch({
        type: T.RECEIVED_HISTORY_FEED,
        offset,
        episodes,
      })
      dispatch({
        type: T.RECEIVED_EPISODE_PLAYBACKS,
        playbacks,
      })

      if (episodes.length < limit) {
        dispatch({
          type: T.RECEIVED_ALL_HISTORY_FEED,
        })
      }
    },
    { requestId: RequestId.getHistoryFeed() },
  )
}
