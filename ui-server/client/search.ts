import { Podcast } from 'types/app'
import * as unmarshal from 'utils/entities'
import { doFetch } from './fetch'

export async function searchPodcasts(
  searchQuery: string,
): Promise<{ podcasts: Podcast[] }> {
  const { data } = await doFetch({
    method: 'GET',
    urlPath: `/results?search_query=${searchQuery}`,
  })

  return { podcasts: (data.results || []).map(unmarshal.podcast) }
}
