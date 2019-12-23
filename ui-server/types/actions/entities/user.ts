import { User } from 'types/app'

export const USER_ADD = 'user/add'

interface AddAction {
  type: typeof USER_ADD
  users: User[]
}

export type UserActionTypes =
  | AddAction
