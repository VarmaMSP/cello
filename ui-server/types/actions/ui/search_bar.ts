export const SEARCH_BAR_UPDATE_TEXT = 'search_bar/update_text'
export const SEARCH_BAR_EXPAND = 'search_bar/expand'
export const SEARCH_BAR_COLLAPSE = 'search_bar/collapse'
export const SEARCH_BAR_UPDATE_TEXT_SUGGESTIONS =
  'search_bar/update_text_suggestions'

export interface UpdateTextAction {
  type: typeof SEARCH_BAR_UPDATE_TEXT
  text: string
}

export interface ExpandAction {
  type: typeof SEARCH_BAR_EXPAND
}

export interface CollapseAction {
  type: typeof SEARCH_BAR_COLLAPSE
}

export interface UpdateTextSuggestionsAction {
  type: typeof SEARCH_BAR_UPDATE_TEXT_SUGGESTIONS
  suggestions: string[]
}

export type SearchBarActionTypes =
  | UpdateTextAction
  | ExpandAction
  | CollapseAction
  | UpdateTextSuggestionsAction
