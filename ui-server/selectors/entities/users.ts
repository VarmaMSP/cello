import { AppState } from 'store'

export function getUserById(state: AppState, userId: string) {
  return state.entities.users.byId[userId]
}
