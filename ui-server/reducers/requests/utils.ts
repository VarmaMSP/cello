import { Reducer } from 'redux'

export type RequestStatus = 'NOT_STARTED' | 'STARTED' | 'SUCCESS' | 'FAILURE'

export interface RequestState {
  status: RequestStatus
  error: string | null
}

export function initialRequestState(): RequestState {
  return {
    status: 'NOT_STARTED',
    error: null,
  }
}

export function defaultRequestReducer(
  requestType: string,
  successType: string,
  failureType: string,
): Reducer<RequestState, { type: string; error: string | null }> {
  return (state = initialRequestState(), action) => {
    switch (action.type) {
      case requestType:
        return { status: 'STARTED', error: null }
      case successType:
        return { status: 'SUCCESS', error: null }
      case failureType:
        return { status: 'FAILURE', error: action.error }
      default:
        return state
    }
  }
}
