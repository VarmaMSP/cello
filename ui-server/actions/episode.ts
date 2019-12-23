import * as client from 'client/episode'
import * as T from 'types/actions'
import { requestAction } from './utils'

export function getEpisodePageData(episodeUrlParam: string) {
  return requestAction(
    () => client.getEpisodePageData(episodeUrlParam),
    (dispatch, _, { podcast, episode }) => {
      dispatch({ type: T.EPISODE_ADD, episodes: [episode] })
      dispatch({ type: T.PODCAST_ADD, podcasts: [podcast] })
    },
  )
}
