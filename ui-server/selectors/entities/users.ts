import { createSelector } from 'reselect'
import { AppState } from 'store'
import { EpisodePlayback, Podcast, User } from 'types/app'
import { $Id, MapById, MapOneToMany } from 'types/utilities'

export function getIsUserSignedIn(state: AppState) {
  return !!state.entities.user.currentUserId
}

export function getIsUserSubscribedToPodcast(
  state: AppState,
  podcastId: string,
) {
  return makeGetUserSubscriptions()(state).some(
    (podcast) => podcast.id === podcastId,
  )
}

export function getUserEpisodePlaybacks(
  state: AppState,
): { [episodeId: string]: EpisodePlayback } {
  return state.entities.user.playback
}

export function makeGetUserSubscriptions() {
  return createSelector<
    AppState,
    $Id<User>,
    MapOneToMany<User, Podcast>,
    MapById<Podcast>,
    Podcast[]
  >(
    (state) => state.entities.user.currentUserId,
    (state) => state.entities.podcasts.podcastsSubscribedByUser,
    (state) => state.entities.podcasts.podcasts,
    (userId, subscriptions, podcasts) => {
      return (subscriptions[userId] || []).map(
        (podcastId) => podcasts[podcastId],
      )
    },
  )
}

export function getCurrenUser() {
  return createSelector<AppState, $Id<User>, MapById<User>, User>(
    (state) => state.entities.user.currentUserId,
    (state) => state.entities.user.users,
    (userId, users) => users[userId],
  )
}
