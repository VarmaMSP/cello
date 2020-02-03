import { Episode, Podcast } from 'types/app'
import { ApiResponse } from 'types/models/api_response'
import * as unmarshal from 'utils/entities'
import { qs } from 'utils/utils'
import { doFetch } from './fetch'

export async function getPodcastPageData(
  podcastUrlParam: string,
): Promise<{ podcasts: Podcast[]; episodes: Episode[] }> {
  const { data } = await doFetch({
    method: 'GET',
    urlPath: `/podcasts/${podcastUrlParam}`,
  })

  const c = new ApiResponse(data)
  return {
    "podcasts": c.podcasts,
    "episodes": c.episodes,
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
