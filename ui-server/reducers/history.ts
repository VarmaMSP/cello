import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { HistoryEntry } from 'types/browser'

// stack[0] should always represent the last page in the session
const stack: Reducer<HistoryEntry[], T.AppActions> = (state = [], action) => {
  switch (action.type) {
    case T.HISTORY_PUSH_ENTRY:
      return [action.entry, ...state]

    case T.HISTORY_POP_ENTRY:
      const [, ...entries] = state
      return entries

    default:
      return state
  }
}

const poppedEntry: Reducer<HistoryEntry, T.AppActions> = (
  state = { urlPath: '', scrollY: 0 },
  action,
) => {
  switch (action.type) {
    case T.HISTORY_POP_ENTRY:
      return action.entry

    default:
      return state
  }
}

export default combineReducers({
  stack,
  poppedEntry,
})
