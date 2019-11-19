export const REQUEST_IN_PROGRESS = 'REQUEST_IN_PROGRESS'
export const REQUEST_COMPLETE = 'REQUEST_COMPLETE'

export interface RequestInProgressAction {
  type: typeof REQUEST_IN_PROGRESS
  requestId: string
}

export interface RequestCompleteAction {
  type: typeof REQUEST_COMPLETE
  requestId: string
}

export type RequestActionTypes = RequestInProgressAction | RequestCompleteAction
