import { AppState } from 'store'
import { Category } from 'types/models'

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
