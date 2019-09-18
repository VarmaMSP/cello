import { RequestException } from 'client/client'
import { Dispatch } from 'redux'
import { AppActions } from 'types/actions'

type ResolveData<T> = T extends Promise<infer U> ? U : T

export function requestAction<T extends Promise<any>>(
  getData: () => T,
  processData: (dispatch: Dispatch, data: ResolveData<T>) => void,
  onRequestAction: AppActions,
  onSuccessAction: AppActions,
  onFailureAction: AppActions,
) {
  return async (dispatch: Dispatch<AppActions>) => {
    dispatch(onRequestAction)
    try {
      const res = await getData()
      processData(dispatch, res)
      dispatch(onSuccessAction)
    } catch (err) {
      err = err as RequestException
      dispatch(onFailureAction)
    }
  }
}
