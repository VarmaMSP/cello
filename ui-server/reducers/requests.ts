import { Reducer } from 'redux'
import * as T from 'types/actions'

const requests: Reducer<
  { [requestId: string]: 'IN_PROGRESS' | 'COMPLETE' },
  T.AppActions
> = (state = {}, action) => {
  switch (action.type) {
    case T.REQUEST_IN_PROGRESS:
      return { ...state, [action.requestId]: 'IN_PROGRESS' }
    case T.REQUEST_COMPLETE:
      return { ...state, [action.requestId]: 'COMPLETE' }
    default:
      return state
  }
}

export default requests
