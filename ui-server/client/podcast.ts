import { Episode, Podcast } from 'types/app'
import * as unmarshal from 'utils/entities'
import { qs } from 'utils/utils'
import { doFetch } from './fetch'

export async function getPodcastPageData(
  podcastUrlParam: string,
): Promise<{ podcast: Podcast; episodes: Episode[] }> {
  const { data } = await doFetch({
    method: 'GET',
    urlPath: `/podcasts/${podcastUrlParam}`,
  })

  return {
    podcast: unmarshal.podcast(data.podcast),
    episodes: (data.episodes || []).map(unmarshal.episode),
  }
}

export async function getPodcastEpisodes(
  podcastId: string,
  limit: number,
  offset: number,
  order: string,
): Promise<{
  episodes: Episode[]
}> {
  const { data } = await doFetch({
    method: 'GET',
    urlPath: `/ajax/browse?${qs({
      endpoint: 'podcast_episodes',
      podcast_id: podcastId,
      order: order,
      offset: offset,
      limit: limit,
    })}`,
  })

  return { episodes: (data.episodes || []).map(unmarshal.episode) }
}
