import { RequestException } from 'client/client'
import { Dispatch } from 'redux'
import { getIsUserSignedIn } from 'selectors/entities/users'
import { AppState } from 'store'
import {
  AppActions,
  SHOW_SIGNIN_MODAL,
  SIGN_OUT_USER_FORCEFULLY,
} from 'types/actions'

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
      if (err.statusCode === 401) {
        const userSignedIn = getIsUserSignedIn(getState())
        if (!userSignedIn) {
          dispatch({ type: SHOW_SIGNIN_MODAL })
        }
        dispatch({ type: SIGN_OUT_USER_FORCEFULLY })
      }
      dispatch(onFailureAction)
    }
  }
}
