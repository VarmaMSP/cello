import { combineReducers, Reducer } from 'redux'
import { AppActions, RECEIVED_PODCAST_CURATION } from 'types/actions'
import { Curation } from 'types/app'

const curations: Reducer<{ [curationId: string]: Curation }, AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case RECEIVED_PODCAST_CURATION:
      return { ...state, [action.curation.id]: action.curation }
    default:
      return state
  }
}

export default combineReducers({
  curations,
})
