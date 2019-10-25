import { RequestException } from 'client/client'
import { Dispatch } from 'redux'
import { AppState } from 'store'
import { AppActions, SIGN_OUT_USER_FORCEFULLY } from 'types/actions'

type ResolveData<T> = T extends Promise<infer U> ? U : T

export function requestAction<T extends Promise<any>>(
  getData: (getState: () => AppState) => T,
  processData: (
    dispatch: Dispatch,
    data: ResolveData<T>,
    getState: () => AppState,
  ) => void,
  onRequestAction: AppActions,
  onSuccessAction: AppActions,
  onFailureAction: AppActions,
) {
  return async (dispatch: Dispatch<AppActions>, getState: () => AppState) => {
    dispatch(onRequestAction)
    try {
      const res = await getData(getState)
      processData(dispatch, res, getState)
      dispatch(onSuccessAction)
    } catch (err) {
      if ((err as RequestException).statusCode === 401) {
        dispatch({ type: SIGN_OUT_USER_FORCEFULLY })
      }
      dispatch(onFailureAction)
    }
  }
}
