import { doFetch_ } from 'client/fetch'
import * as T from 'types/actions'
import { requestAction } from './utils'

export function getEpisodePageData(episodeUrlParam: string) {
  return requestAction(
    () =>
      doFetch_({
        method: 'GET',
        urlPath: `/episodes/${episodeUrlParam}`,
      }),
    (dispatch, _, { podcasts, episodes }) => {
      dispatch({ type: T.EPISODE_ADD, episodes: episodes })
      dispatch({ type: T.PODCAST_ADD, podcasts: podcasts })
    },
  )
}
