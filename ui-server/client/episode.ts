import { Episode, Podcast } from 'types/app'
import * as unmarshal from 'utils/entities'
import { doFetch } from './fetch'

export async function getEpisodePageData(
  episodeUrlParam: string,
): Promise<{ podcast: Podcast; episode: Episode }> {
  const { data } = await doFetch({
    method: 'GET',
    urlPath: `/episodes/${episodeUrlParam}`,
  })

  return {
    podcast: unmarshal.podcast(data.podcast),
    episode: unmarshal.episode(data.episode),
  }
}
