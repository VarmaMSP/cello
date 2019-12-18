import { Episode } from 'types/app'
import * as unmarshal from 'utils/entities'
import { qs } from 'utils/utils'
import { doFetch } from './fetch'

export async function getSubscriptionsPageData(): Promise<{
  episodes: Episode[]
}> {
  const { data } = await doFetch({
    method: 'GET',
    urlPath: `/subscriptions`,
  })

  return { episodes: (data.episodes || []).map(unmarshal.episode) }
}

export async function getSubscriptionsFeed(
  offset: number,
  limit: number,
): Promise<{ episodes: Episode[] }> {
  const { data } = await doFetch({
    method: 'GET',
    urlPath: `/ajax/browse?${qs({
      endpoint: 'subscriptions_feed',
      offset: offset,
      limit: limit,
    })}`,
  })

  return { episodes: (data.episodes || []).map(unmarshal.episode) }
}

export async function subscribePodcast(podcastId: string): Promise<void> {
  await doFetch({
    method: 'POST',
    urlPath: `/ajax/service?${qs({
      endpoint: 'subscribe_podcast',
    })}`,
    body: {
      podcast_id: podcastId,
    },
  })
}

export async function unsubscribePodcast(podcastId: string): Promise<void> {
  await doFetch({
    method: 'POST',
    urlPath: `/ajax/service?${qs({
      endpoint: 'unsubscribe_podcast',
    })}`,
    body: {
      podcast_id: podcastId,
    },
  })
}
