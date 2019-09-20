import { RequestException } from 'client/client'
import { Dispatch } from 'redux'
import { AppState } from 'store'
import { AppActions } from 'types/actions'

type ResolveData<T> = T extends Promise<infer U> ? U : T

export function requestAction<T extends Promise<any>>(
  getData: () => T,
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
      const res = await getData()
      processData(dispatch, res, getState)
      dispatch(onSuccessAction)
    } catch (err) {
      err = err as RequestException
      dispatch(onFailureAction)
    }
  }
}
