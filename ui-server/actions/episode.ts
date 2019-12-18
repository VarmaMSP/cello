import * as client from 'client/episode'
import * as T from 'types/actions'
import { requestAction } from './utils'

export function getEpisodePageData(episodeUrlParam: string) {
  return requestAction(
    () => client.getEpisodePageData(episodeUrlParam),
    (dispatch, _, { podcast, episode }) => {
      dispatch({ type: T.RECEIVED_EPISODE, episode })
      dispatch({ type: T.RECEIVED_PODCAST, podcast })
    },
  )
}
