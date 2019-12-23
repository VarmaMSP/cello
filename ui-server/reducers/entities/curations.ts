import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { Curation, CurationType } from 'types/app'
import { addKeyToArr } from 'utils/immutable'

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

    default:
      return state
  }
}

const byType: Reducer<{ [key in CurationType]: string[] }, T.AppActions> = (
  state = { CATEGORY: [], NORMAL: [] },
  action,
) => {
  switch (action.type) {
    case T.CURATION_ADD:
      return {
        ...state,
        ...action.curations.reduce<{ [key in CurationType]: string[] }>(
          (acc, c) => ({
            ...acc,
            [c.type]: addKeyToArr(c.id, state[c.type]),
          }),
          { CATEGORY: [], NORMAL: [] },
        ),
      }

    default:
      return state
  }
}

export default combineReducers({
  byId,
  byType,
})
