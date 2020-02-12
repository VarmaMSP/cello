import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { Curation } from 'models'

const byId: Reducer<{ [curationId: string]: Curation }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.CURATION_ADD:
      return action.curations.reduce<{ [curationId: string]: Curation }>(
        (acc, c) => ({ ...acc, [c.id]: { ...(state[c.id] || {}), ...c } }),
        state,
      )

    case T.CURATION_ADD_PODCASTS:
      return {
        ...state,
        [action.curationId]: {
          ...(state[action.curationId] || {}),
          podcastIds: action.podcastIds,
        },
      }

    default:
      return state
  }
}

export default combineReducers({
  byId,
})
