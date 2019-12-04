import * as client from 'client/playback'
import { Dispatch } from 'redux'
import { getIsUserSignedIn } from 'selectors/entities/users'
import { getPlayingEpisodeId } from 'selectors/ui/player'
import { AppState } from 'store'
import * as T from 'types/actions'

export function getEpisodePlaybacks(episodeIds: string[]) {
  return async (dispatch: Dispatch<T.AppActions>, getState: () => AppState) => {
    if (!getIsUserSignedIn(getState())) {
      return
    }

    try {
      const { playbacks } = await client.getPlaybacks(episodeIds)
      dispatch({ type: T.RECEIVED_PLAYBACKS, playbacks })
    } catch (err) {
      console.log(err)
    }
  }
}

export function startPlayback(episodeId: string, currentTime: number) {
  return async (dispatch: Dispatch<T.AppActions>, getState: () => AppState) => {
    const state = getState()
    if (!getIsUserSignedIn(state) || getPlayingEpisodeId(state) === episodeId) {
      dispatch({ type: T.PLAY_EPISODE, episodeId, currentTime })
      return
    }

    try {
      await client.startPlayback(episodeId)
      dispatch({ type: T.PLAY_EPISODE, episodeId, currentTime })
    } catch (err) {
      console.log(err)
    }
  }
}

export function syncPlayback(episodeId: string) {
  return async (_: Dispatch<T.AppActions>, getState: () => AppState) => {
    if (episodeId.length === 0) {
      // THIS WILL PREVENT WRONG API CALLS
      return
    }

    try {
      if (getIsUserSignedIn(getState())) {
        await client.syncPlayback(episodeId)
      }
    } catch (err) {
      console.log(err)
    }
  }
}
