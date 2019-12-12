import { FetchException } from 'client/fetch'
import { Dispatch } from 'redux'
import { getCurrentUserId } from 'selectors/entities/users'
import { requestStatus } from 'selectors/request'
import { AppState } from 'store'
import * as AT from 'types/actions'

type MakeRequest<T> = (g: () => AppState) => T

type ProcessData<T> = (
  d: Dispatch<AT.AppActions>,
  g: () => AppState,
  r: T extends Promise<infer U> ? U : T,
) => void

type RequestActionOpts = {
  requestId: string
  skip: RequestActionSkipCond
  notifyError: boolean
}

type RequestActionSkipCond =
  | { cond: 'USER_NOT_SIGNED_IN' }
  | { cond: 'REQUEST_ALREADY_MADE' }

export function requestAction<T extends Promise<any>>(
  makeRequest: MakeRequest<T>,
  processData: ProcessData<T>,
  { skip, requestId }: Partial<RequestActionOpts> = {},
) {
  return async (
    dispatch: Dispatch<AT.AppActions>,
    getState: () => AppState,
  ) => {
    if (!!skip) {
      switch (skip.cond) {
        case 'REQUEST_ALREADY_MADE':
          if (requestStatus(getState(), requestId!) === 'SUCCESS') {
            return
          }
          break
        case 'USER_NOT_SIGNED_IN':
          if (getCurrentUserId(getState()) === '') {
            return
          }
          break
      }
    }

    !!requestId && dispatch({ type: AT.REQUEST_IN_PROGRESS, requestId })

    try {
      const res = await makeRequest(getState)
      processData(dispatch, getState, res)
      !!requestId && dispatch({ type: AT.REQUEST_SUCCESS, requestId })
    } catch (err) {
      if ((err as FetchException).statusCode === 401) {
        dispatch({ type: AT.SIGN_OUT_USER_FORCEFULLY })
      }
      !!requestId && dispatch({ type: AT.REQUEST_FAILURE, requestId })
    }
  }
}
