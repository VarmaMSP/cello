import { doFetch } from 'utils/fetch'
import * as T from 'types/actions'
import { requestAction } from './utils'

export function getEpisodePageData(episodeUrlParam: string) {
  return requestAction(
    () =>
      doFetch({
        method: 'GET',
        urlPath: `/episodes/${episodeUrlParam}`,
      }),
    (dispatch, _, { podcasts, episodes }) => {
      dispatch({ type: T.EPISODE_ADD, episodes: episodes })
      dispatch({ type: T.PODCAST_ADD, podcasts: podcasts })
    },
  )
}
