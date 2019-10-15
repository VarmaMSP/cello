import client from 'client'
import { getPlayingEpisodeId } from 'selectors/ui/player'
import * as T from 'types/actions'
import { requestAction } from './utils'

export function playEpisode(episodeId: string) {
  return requestAction(
    (getState) =>
      getPlayingEpisodeId(getState()) === episodeId
        ? Promise.resolve()
        : client.syncPlayback(episodeId),
    (dispatch) => {
      dispatch({ type: T.PLAY_EPISODE, episodeId: episodeId })
    },
    { type: T.SYNC_PLAYBACK_REQUEST },
    { type: T.SYNC_PLAYBACK_SUCCESS },
  )
}
