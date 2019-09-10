export type RequestStatus = 'NOT_STARTED' | 'STARTED' | 'SUCCESS' | 'FAILURE'

export interface RequestState {
  status: RequestStatus
  error: Error | null
}

export function initialRequestState(): RequestState {
  return {
    status: 'NOT_STARTED',
    error: null,
  }
}
