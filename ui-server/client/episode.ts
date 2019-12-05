import { Episode, Podcast } from 'types/app'
import * as unmarshal from 'utils/entities'
import { doFetch } from './fetch'

export async function getEpisode(
  episodeId: string,
): Promise<{ podcast: Podcast; episode: Episode }> {
  const { data } = await doFetch({
    method: 'GET',
    urlPath: `/episodes/${episodeId}`,
  })

  return {
    podcast: unmarshal.podcast(data.podcast),
    episode: unmarshal.episode(data.episode),
  }
}

export async function getPodcastEpisodes(
  podcastId: string,
  limit: number,
  offset: number,
  order: 'pub_date_desc' | 'pub_date_asc',
): Promise<{
  episodes: Episode[]
}> {
  const { data } = await doFetch({
    method: 'GET',
    urlPath: `/podcasts/${podcastId}/episodes?limit=${limit}&offset=${offset}&order=${order}`,
  })

  return { episodes: (data.episodes || []).map(unmarshal.episode) }
}
