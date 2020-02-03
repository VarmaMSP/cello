import { doFetch } from 'utils/fetch'
import { Dispatch } from 'redux'
import { getEpisodeById } from 'selectors/entities/episodes'
import { getIsUserSignedIn } from 'selectors/session'
import { getPlayingEpisodeId } from 'selectors/ui/audio_player'
import { AppState } from 'store'
import * as T from 'types/actions'
import * as gtag from 'utils/gtag'
import { requestAction } from './utils'

export function getEpisodePlaybacks(episodeIds: string[]) {
  return requestAction(
    () =>
      doFetch({
        method: 'POST',
        urlPath: `/ajax/service?endpoint=get_playbacks`,
        body: { episode_ids: episodeIds },
      }),
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
    gtag.playEpisode(
      (getEpisodeById(getState(), episodeId) || {}).title,
      beginAt,
    )

    try {
      const state = getState()
      if (
        getIsUserSignedIn(state) &&
        getPlayingEpisodeId(state) !== episodeId
      ) {
        await doFetch({
          method: 'POST',
          urlPath: `/ajax/service?endpoint=playback_sync&action=playback_begin`,
          body: { episode_id: episodeId },
        })
      }
      if (getPlayingEpisodeId(state) !== episodeId) {
        dispatch({ type: T.AUDIO_PLAYER_PLAY_EPISODE, episodeId, beginAt })
      }
    } catch (err) {}
  }
}

export function syncPlayback(episodeId: string, position: number) {
  return requestAction(
    () =>
      doFetch({
        method: 'POST',
        urlPath: `/ajax/service?endpoint=playback_sync&action=playback_progress`,
        body: { episode_id: episodeId, position: Number(position.toFixed(6)) },
      }),
    () => {},
    { skip: { cond: 'USER_NOT_SIGNED_IN' } },
  )
}
