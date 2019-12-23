export const REQUEST_IN_PROGRESS = 'request/in_progress'
export const REQUEST_SUCCESS = 'request/success'
export const REQUEST_FAILURE = 'request/failure'

interface InProgressAction {
  type: typeof REQUEST_IN_PROGRESS
  requestId: string
}

interface SuccessAction {
  type: typeof REQUEST_SUCCESS
  requestId: string
}

interface FailureAction {
  type: typeof REQUEST_FAILURE
  requestId: string
}

export type RequestActionTypes =
  | InProgressAction
  | SuccessAction
  | FailureAction
