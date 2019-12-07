export const REQUEST_IN_PROGRESS = 'REQUEST_IN_PROGRESS'
export const REQUEST_SUCCESS = 'REQUEST_SUCCESS'
export const REQUEST_FAILURE = 'REQUEST_FAILURE'

export interface RequestInProgressAction {
  type: typeof REQUEST_IN_PROGRESS
  requestId: string
}

export interface RequestSuccessAction {
  type: typeof REQUEST_SUCCESS
  requestId: string
}

export interface RequestFailureAction {
  type: typeof REQUEST_FAILURE
  requestId: string
}

export type RequestActionTypes =
  | RequestInProgressAction
  | RequestSuccessAction
  | RequestFailureAction
