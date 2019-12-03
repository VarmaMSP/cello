import { Episode } from 'types/app'
import * as unmarshal from 'utils/entities'
import { doFetch } from './fetch'

export async function getSubscriptionsFeed(
  offset: number,
  limit: number,
): Promise<{ episodes: Episode[] }> {
  const { data } = await doFetch({
    method: 'GET',
    urlPath: `/subscriptions/feed?limit=${limit}&offset=${offset}`,
  })

  return { episodes: (data.episodes || []).map(unmarshal.episode) }
}

export async function subscribeToPodcast(podcastId: string): Promise<void> {
  await doFetch({
    method: 'PUT',
    urlPath: `/podcasts/${podcastId}/subscribe`,
  })
}

export async function unsubscribeToPodcast(podcastId: string): Promise<void> {
  await doFetch({
    method: 'PUT',
    urlPath: `/podcasts/${podcastId}/unsubscribe`,
  })
}
