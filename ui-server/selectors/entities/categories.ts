import { Category } from 'models'
import { createSelector } from 'reselect'
import { AppState } from 'store'
import { MapById } from 'types/utilities'

export function getCategoryById(state: AppState, categoryId: string) {
  return state.entities.categories.byId[categoryId]
}

export function getCategoriesByIds(
  state: AppState,
  categoryIds: string[],
): Category[] {
  return categoryIds.map((x) => state.entities.categories.byId[x])
}

export function getAllCategories(state: AppState) {
  return Object.keys(state.entities.categories.byId).map(
    (x) => state.entities.categories.byId[x],
  )
}

export function makeGetPrimaryCategories() {
  return createSelector<AppState, MapById<Category>, Category[]>(
    (state) => state.entities.categories.byId,
    (all) =>
      Object.keys(all)
        .map((x) => all[x])
        .filter((x) => !!!x.parentId),
  )
}
