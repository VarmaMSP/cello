import { Reducer } from 'redux'
import * as T from 'types/actions'

type RequestStatus = 'NOT_STARTED' | 'IN_PROGRESS' | 'SUCCESS' | 'FAILURE'

const requests: Reducer<
  {[requestId: string]: RequestStatus},
  T.AppActions
> = (state = {}, action) => {
  switch (action.type) {
    case T.REQUEST_IN_PROGRESS:
      return { ...state, [action.requestId]: 'IN_PROGRESS' }
    case T.REQUEST_SUCCESS:
      return { ...state, [action.requestId]: 'SUCCESS' }
    case T.REQUEST_FAILURE:
      return { ...state, [action.requestId]: 'FAILURE' }
    default:
      return state
  }
}

export default requests
