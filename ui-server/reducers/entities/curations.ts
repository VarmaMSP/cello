import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { Curation } from 'types/models'

const byId: Reducer<{ [curationId: string]: Curation }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.CURATION_ADD:
      return {
        ...state,
        ...action.curations.reduce<{ [curationId: string]: Curation }>(
          (acc, c) => ({ ...acc, [c.id]: { ...(state[c.id] || {}), ...c } }),
          {},
        ),
      }

    case T.CURATION_ADD_PODCASTS:
      return {
        ...state,
        [action.curationId]: {
          ...(state[action.curationId] || {}),
          members: action.podcastIds,
        },
      }

    default:
      return state
  }
}

const byType: Reducer<{ [key: string]: string[] }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.CURATION_ADD:
      return {
        ...state,
        [action.curationType]: action.curations.map((x) => x.id),
      }

    default:
      return state
  }
}

export default combineReducers({
  byId,
  byType,
})
