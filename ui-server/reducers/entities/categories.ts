import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { Category } from 'types/models'

const byId: Reducer<{ [categoryId: string]: Category }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.CATEGORY_ADD:
      return action.categories.reduce<{ [categoryId: string]: Category }>(
        (acc, c) => ({ ...acc, [c.id]: { ...(acc[c.id] || {}), ...c } }),
        state,
      )

    default:
      return state
  }
}

export default combineReducers({
  byId,
})
