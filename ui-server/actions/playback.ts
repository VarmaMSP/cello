import * as client from 'client/playback'
import { getIsUserSignedIn } from 'selectors/session'
import { getPlayingEpisodeId } from 'selectors/ui/audio_player'
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
  return requestAction(
    () => client.startPlayback(episodeId),
    (dispatch) => {
      dispatch({ type: T.AUDIO_PLAYER_PLAY_EPISODE, episodeId, beginAt })
    },
    {
      skip: {
        cond: 'CUSTOM',
        p: (getState) =>
          !getIsUserSignedIn(getState()) ||
          getPlayingEpisodeId(getState()) == '',
      },
    },
  )
}

export function syncPlayback(episodeId: string) {
  return requestAction(
    () => client.syncPlayback(episodeId),
    () => {},
    { skip: { cond: 'USER_NOT_SIGNED_IN' } },
  )
}
