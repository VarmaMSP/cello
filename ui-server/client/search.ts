import {
  Episode,
  EpisodeSearchResult,
  Podcast,
  PodcastSearchResult,
} from 'types/app'
import * as unmarshal from 'utils/entities'
import { qs } from 'utils/utils'
import { doFetch } from './fetch'

export async function getResultsPageData(
  searchQuery: string,
  resultType: 'podcast' | 'episode',
  sortBy: 'relevance' | 'publish_date',
): Promise<{
  podcasts: Podcast[]
  episodes: Episode[]
  podcastSearchResults: PodcastSearchResult[]
  episodeSearchResults: EpisodeSearchResult[]
}> {
  console.log('something is working')
  const { data } = await doFetch({
    method: 'GET',
    urlPath: `/results?query=${qs({
      query: searchQuery,
      type: resultType,
      sort_by: sortBy,
    })}`,
  })

  return {
    podcasts: (data.podcasts || []).map(unmarshal.podcast),
    episodes: (data.episodes || []).map(unmarshal.episode),
    podcastSearchResults: (data.podcast_search_results || []).map(
      unmarshal.podcastSearchResult,
    ),
    episodeSearchResults: (data.episode_search_results || []).map(
      unmarshal.episodeSearchResult,
    ),
  }
}

export async function getResults(
  searchQuery: string,
  resultType: 'podcast' | 'episode',
  sortBy: 'relevance' | 'publish_date',
  offset: number,
  limit: number,
): Promise<{
  podcasts: Podcast[]
  episodes: Episode[]
  podcastSearchResults: PodcastSearchResult[]
  episodeSearchResults: EpisodeSearchResult[]
}> {
  const { data } = await doFetch({
    method: 'GET',
    urlPath: `/ajax/browse?${qs({
      endpoint: 'search_results',
      query: searchQuery,
      type: resultType,
      sort_by: sortBy,
      offset: offset,
      limit: limit,
    })}`,
  })

  return {
    podcasts: (data.podcasts || []).map(unmarshal.podcast),
    episodes: (data.episodes || []).map(unmarshal.episode),
    podcastSearchResults: (data.podcast_search_results || []).map(
      unmarshal.podcastSearchResult,
    ),
    episodeSearchResults: (data.episode_search_results || []).map(
      unmarshal.episodeSearchResult,
    ),
  }
}
