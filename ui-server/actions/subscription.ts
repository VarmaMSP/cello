import { doFetch } from 'utils/fetch'
import { getPodcastById } from 'selectors/entities/podcasts'
import * as T from 'types/actions'
import * as gtag from 'utils/gtag'
import * as RequestId from 'utils/request_id'
import { qs } from 'utils/utils'
import { requestAction } from './utils'

export function getSubscriptionsPageData() {
  return requestAction(
    () =>
      doFetch({
        method: 'GET',
        urlPath: `/subscriptions`,
      }),
    (dispatch, _, { episodes }) => {
      dispatch({ type: T.EPISODE_ADD, episodes })

      dispatch({
        type: T.SUBSCRIPTIONS_FEED_LOAD_PAGE,
        page: 0,
        episodeIds: episodes.map((x) => x.id),
      })

      if (episodes.length < 15) {
        dispatch({ type: T.SUBSCRIPTIONS_FEED_RECEIVED_ALL })
      }
    },
    { requestId: RequestId.getSubscriptionsPageData() },
  )
}

export function getSubscriptionsFeed(offset: number, limit: number) {
  return requestAction(
    () =>
      doFetch({
        method: 'GET',
        urlPath: `/ajax/browse?${qs({
          endpoint: 'subscriptions_feed',
          offset: offset,
          limit: limit,
        })}`,
      }),
    (dispatch, _, { episodes }) => {
      dispatch({ type: T.EPISODE_ADD, episodes })

      dispatch({
        type: T.SUBSCRIPTIONS_FEED_LOAD_PAGE,
        page: Math.floor(offset / 10),
        episodeIds: episodes.map((x) => x.id),
      })

      if (episodes.length < limit) {
        dispatch({ type: T.SUBSCRIPTIONS_FEED_RECEIVED_ALL })
      }
    },
    { requestId: RequestId.getSubscriptionsFeed() },
  )
}

export function subscribeToPodcast(podcastId: string) {
  return requestAction(
    () =>
      doFetch({
        method: 'POST',
        urlPath: `/ajax/service?${qs({
          endpoint: 'subscribe_podcast',
        })}`,
        body: {
          podcast_id: podcastId,
        },
      }),
    (dispatch, getState) => {
      gtag.subscribePodcast((getPodcastById(getState(), podcastId) || {}).title)

      dispatch({
        type: T.SESSION_SUBSCRIBE_PODCASTS,
        podcastIds: [podcastId],
      })
    },
  )
}

export function unsubscribeToPodcast(podcastId: string) {
  return requestAction(
    () =>
      doFetch({
        method: 'POST',
        urlPath: `/ajax/service?${qs({
          endpoint: 'unsubscribe_podcast',
        })}`,
        body: {
          podcast_id: podcastId,
        },
      }),
    (dispatch) => {
      dispatch({
        type: T.SESSION_UNSUBSCRIBE_PODCASTS,
        podcastIds: [podcastId],
      })
    },
  )
}
