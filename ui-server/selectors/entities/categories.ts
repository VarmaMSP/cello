import { AppState } from 'store'
import { Category } from 'types/models'

export function getCategoriesByIds(
  state: AppState,
  categoryIds: string[],
): Category[] {
  return categoryIds.map((x) => state.entities.categories.byId[x])
}
