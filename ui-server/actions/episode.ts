import client from 'client'
import { Dispatch } from 'redux'
import { getIsUserSignedIn } from 'selectors/entities/users'
import { getPlayingEpisodeId } from 'selectors/ui/player'
import { AppState } from 'store'
import * as T from 'types/actions'
import { requestAction } from './utils'

export function beginPlayback(episodeId: string, currentTime: number) {
  return async (dispatch: Dispatch<T.AppActions>, getState: () => AppState) => {
    const state = getState()
    if (!getIsUserSignedIn(state) || getPlayingEpisodeId(state) === episodeId) {
      dispatch({ type: T.PLAY_EPISODE, episodeId, currentTime })
      return
    }

    try {
      await client.syncPlayback(episodeId)
      dispatch({ type: T.PLAY_EPISODE, episodeId, currentTime })
    } catch (err) {}
  }
}

export function syncPlayback(episodeId: string, currentTime: number) {
  return async (_: Dispatch<T.AppActions>, getState: () => AppState) => {
    if (episodeId.length === 0) {
      // THIS WILL PREVENT WRONG API CALLS
      return
    }

    if (!getIsUserSignedIn(getState())) {
      return
    }

    try {
      await client.syncPlaybackProgress(episodeId, currentTime)
    } catch (err) {}
  }
}

export function getEpisode(episodeId: string) {
  return requestAction(
    () => client.getEpisode(episodeId),
    (dispatch, _, { podcast, episode }) => {
      dispatch({
        type: T.RECEIVED_EPISODES,
        episodes: [episode],
        podcastId: podcast.id,
      })
      dispatch({ type: T.RECEIVED_PODCAST, podcast })
    },
  )
}

export function getEpisodePlaybacks(episodeIds: string[]) {
  return async (dispatch: Dispatch<T.AppActions>, getState: () => AppState) => {
    if (!getIsUserSignedIn(getState())) {
      return
    }

    try {
      const { playbacks } = await client.getEpisodePlaybacks(episodeIds)
      dispatch({
        type: T.RECEIVED_EPISODE_PLAYBACKS,
        playbacks,
      })
    } catch (err) {
      console.log(err)
    }
  }
}
