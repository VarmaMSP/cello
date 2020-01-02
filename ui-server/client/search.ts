import { Podcast } from 'types/app'
import * as unmarshal from 'utils/entities'
import { qs } from 'utils/utils'
import { doFetch } from './fetch'

export async function getResultsPageData(
  searchQuery: string,
): Promise<{ podcasts: Podcast[] }> {
  const { data } = await doFetch({
    method: 'GET',
    urlPath: `/results?query=${searchQuery}`,
  })

  return { podcasts: (data.results || []).map(unmarshal.podcast) }
}

export async function getResults(
  searchQuery: string,
  offset: number,
  limit: number,
): Promise<{ podcasts: Podcast[] }> {
  const { data } = await doFetch({
    method: 'GET',
    urlPath: `/ajax/browse?${qs({
      endpoint: 'search_results',
      query: searchQuery,
      offset,
      limit,
    })}`,
  })

  return { podcasts: (data.results || []).map(unmarshal.podcast) }
}
