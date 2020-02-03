import { Episode, Podcast } from 'types/app'
import * as unmarshal from 'utils/entities'
import { qs } from 'utils/utils'
import { doFetch } from './fetch'

export async function getHistoryPageData(): Promise<{
  podcasts: Podcast[]
  episodes: Episode[]
}> {
  const { data } = await doFetch({
    method: 'GET',
    urlPath: '/history',
  }) 

  return {
    podcasts: (data.podcasts || []).map(unmarshal.podcast),
    episodes: (data.episodes || []).map(unmarshal.episode),
  }
}

export async function getHistoryFeed(
  offset: number,
  limit: number,
): Promise<{
  podcasts: Podcast[]
  episodes: Episode[]
}> {
  const { data } = await doFetch({
    method: 'GET',
    urlPath: `/ajax/browse?${qs({
      endpoint: 'history_feed',
      offset,
      limit,
    })}`,
  })

  return {
    podcasts: (data.podcasts || []).map(unmarshal.podcast),
    episodes: (data.episodes || []).map(unmarshal.episode),
  }
}
