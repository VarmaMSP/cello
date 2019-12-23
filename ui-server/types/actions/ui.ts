// Searchbar Actions
export const SEARCH_BAR_TEXT_CHANGE = 'SEARCH_BAR_TEXT_CHANGE'

export interface SearchBarTextChangeAction {
  type: typeof SEARCH_BAR_TEXT_CHANGE
  text: string
}

export type UiActionTypes = SearchBarTextChangeAction
