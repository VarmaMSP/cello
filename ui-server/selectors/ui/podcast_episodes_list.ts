import { createSelector } from 'reselect'
import { AppState } from 'store'
import { Episode, Podcast } from 'types/models'
import { PodcastEpisodeListOrder } from 'types/ui'
import { $Id } from 'types/utilities'

export function makeSelectPodcastEpisodeList() {
  return createSelector<
    AppState,
    { podcastId: $Id<Podcast>; order: PodcastEpisodeListOrder },
    { [page: string]: $Id<Episode>[] },
    boolean,
    [$Id<Episode>[], boolean]
  >(
    (state, { podcastId, order }) =>
      (state.ui.podcastEpisodeList.list[podcastId] || {})[order] || {},

    (state, { podcastId, order }) =>
      (state.ui.podcastEpisodeList.receivedAll[podcastId] || []).includes(
        order,
      ),

    (obj, receivedAll) => [
      Object.keys(obj).reduce<$Id<Episode>[]>(
        (acc, key) => [...acc, ...obj[key]],
        [],
      ),
      receivedAll,
    ],
  )
}
