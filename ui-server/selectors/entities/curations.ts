import { AppState } from 'store'
import { Curation } from 'types/app'
import { $Id } from 'types/utilities'

export function getAllCurations(state: AppState) {
  return state.entities.curations.curations
}

export function getCurationById(state: AppState, curationId: $Id<Curation>) {
  return getAllCurations(state)[curationId]
}

export function getAllCurationIds(state: AppState) {
  return Object.keys(state.entities.curations.curations)
}
