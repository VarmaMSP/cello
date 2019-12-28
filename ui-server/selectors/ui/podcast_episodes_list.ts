import { createSelector } from 'reselect'
import { AppState } from 'store'
import { Episode } from 'types/app'
import { $Id } from 'types/utilities'

export function makeGetEpisodeIds() {
  return createSelector<AppState, { [page: ]}
}