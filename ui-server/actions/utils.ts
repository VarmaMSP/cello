import { RequestException } from 'client/client'
import { Dispatch } from 'redux'
import { AppState } from 'store'
import {
  AppActions,
  REQUEST_COMPLETE,
  REQUEST_IN_PROGRESS,
  SIGN_OUT_USER_FORCEFULLY,
} from 'types/actions'

interface RequestActionOpts {
  requestId: string
  notifyError: boolean
}

type ResolveData<T> = T extends Promise<infer U> ? U : T

export function requestAction<T extends Promise<any>>(
  makeRequest: () => T,
  processData: (
    dispatch: Dispatch<AppActions>,
    getState: () => AppState,
    data: ResolveData<T>,
  ) => void,
  { requestId }: Partial<RequestActionOpts> = {},
) {
  return async (dispatch: Dispatch<AppActions>, getState: () => AppState) => {
    !!requestId && dispatch({ type: REQUEST_IN_PROGRESS, requestId })
    try {
      const res = await makeRequest()
      processData(dispatch, getState, res)
    } catch (err) {
      if ((err as RequestException).statusCode === 401) {
        dispatch({ type: SIGN_OUT_USER_FORCEFULLY })
      }
    }
    !!requestId && dispatch({ type: REQUEST_COMPLETE, requestId })
  }
}
