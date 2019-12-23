import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { CurationMember } from 'types/app'
import { addKeysToArr } from 'utils/immutable'

function makeId(curationId: string, podcastId: string) {
  return `${curationId}${podcastId}`
}

const byId: Reducer<{ [id: string]: CurationMember }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.CURATION_ADD_PODCASTS:
      return {
        ...state,
        ...action.podcastIds.reduce<{ [id: string]: CurationMember }>(
          (acc, podcastId) => ({
            ...acc,
            [makeId(action.curationId, podcastId)]: {
              id: makeId(action.curationId, podcastId),
              podcastId: podcastId,
              curationId: action.curationId,
            },
          }),
          {},
        ),
      }

    default:
      return state
  }
}

const byCurationId: Reducer<{ [curation: string]: string[] }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.CURATION_ADD_PODCASTS:
      return {
        ...state,
        [action.curationId]: addKeysToArr(
          action.podcastIds.map((podcastId) =>
            makeId(action.curationId, podcastId),
          ),
          state[action.curationId] || [],
        ),
      }

    default:
      return state
  }
}

export default combineReducers({
  byId,
  byCurationId,
})
