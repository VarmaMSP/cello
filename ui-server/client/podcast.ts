import { Episode, Podcast } from 'types/app'
import * as unmarshal from 'utils/entities'
import { doFetch } from './fetch'

export async function getPodcast(
  podcastId: string,
): Promise<{ podcast: Podcast; episodes: Episode[] }> {
  const { data } = await doFetch({
    method: 'GET',
    urlPath: `/podcasts/${podcastId}`,
  })

  return {
    podcast: unmarshal.podcast(data.podcast),
    episodes: (data.episodes || []).map(unmarshal.episode),
  }
}
