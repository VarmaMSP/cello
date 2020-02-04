import { Category } from 'types/models/category'

export const CATEGORY_ADD = 'category_add'

interface AddAction {
  type: typeof CATEGORY_ADD
  categories: Category[]
}

export type CategoryActionTypes = AddAction
