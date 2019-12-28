import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { Curation, CurationType } from 'types/app'
import { addKeysToArr } from 'utils/immutable'

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
        CATEGORY: addKeysToArr(
          action.curations
            .filter((c) => c.type === 'CATEGORY')
            .map((x) => x.id),
          state['CATEGORY'],
        ),
        NORMAL: addKeysToArr(
          action.curations.filter((c) => c.type === 'NORMAL').map((x) => x.id),
          state['NORMAL'],
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
