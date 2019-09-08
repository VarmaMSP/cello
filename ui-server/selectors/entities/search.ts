import { createSelector } from 'reselect'
import { AppState } from 'store'
import { Podcast } from 'types/app'
import { $Id, MapById } from 'types/utilities'
import { getAllPodcasts } from './podcasts'

export function getSearchResults() {
  return createSelector<AppState, MapById<Podcast>, $Id<Podcast>[], Podcast[]>(
    getAllPodcasts,
    (state) => state.entities.search.podcasts,
    (podcasts, ids) => ids.map((id) => podcasts[id]),
  )
}
