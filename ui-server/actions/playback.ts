import * as client from 'client/playback'
import { Dispatch } from 'redux'
import { getIsUserSignedIn } from 'selectors/session'
import { getPlayingEpisodeId } from 'selectors/ui/audio_player'
import { AppState } from 'store'
import * as T from 'types/actions'
import { requestAction } from './utils'

export function getEpisodePlaybacks(episodeIds: string[]) {
  return requestAction(
    () => client.getPlaybacks(episodeIds),
    (dispatch, _, { playbacks }) => {
      dispatch({
        type: T.EPISODE_JOIN_PLAYBACK,
        playbacks,
      })
    },
    { skip: { cond: 'USER_NOT_SIGNED_IN' } },
  )
}

export function startPlayback(episodeId: string, beginAt: number) {
  return async (dispatch: Dispatch<T.AppActions>, getState: () => AppState) => {
    try {
      const state = getState()
      if (getIsUserSignedIn(state) && getPlayingEpisodeId(state) !== episodeId) {
        await client.startPlayback(episodeId)
      }
      if (getPlayingEpisodeId(state) !== episodeId) {
        dispatch({ type: T.AUDIO_PLAYER_PLAY_EPISODE, episodeId, beginAt })
      }
    } catch (err) {}
  }
}

export function syncPlayback(episodeId: string, position: number) {
  return requestAction(
    () => client.syncPlayback(episodeId, position),
    () => {},
    { skip: { cond: 'USER_NOT_SIGNED_IN' } },
  )
}
