import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { Curation } from 'types/app'

const curations: Reducer<{ [curationId: string]: Curation }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.RECEIVED_PODCAST_CURATION:
      return { ...state, [action.curation.id]: action.curation }
    default:
      return state
  }
}

export default combineReducers({
  curations,
})
