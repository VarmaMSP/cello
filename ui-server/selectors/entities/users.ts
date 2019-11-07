import { createSelector } from 'reselect'
import { AppState } from 'store'
import { User } from 'types/app'
import { $Id, MapById } from 'types/utilities'

export function getIsUserSignedIn(state: AppState) {
  return !!state.entities.user.currentUserId
}

export function getCurrentUserId(state: AppState) {
  return state.entities.user.currentUserId
}

export function getCurrenUser() {
  return createSelector<AppState, $Id<User>, MapById<User>, User>(
    (state) => state.entities.user.currentUserId,
    (state) => state.entities.user.users,
    (userId, users) => users[userId],
  )
}
