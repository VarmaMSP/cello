export const SEARCH_BAR_UPDATE_TEXT = 'search_bar/update_text'
export const SEARCH_BAR_EXPAND = 'search_bar/expand'
export const SEARCH_BAR_COLLAPSE = 'search_bar/collapse'
export const SEARCH_BAR_SET_SHOW_SUGGESTIONS = 'search_bar/show_suggestions'
export const SEARCH_BAR_UPDATE_TEXT_SUGGESTIONS =
  'search_bar/update_text_suggestions'

export interface UpdateTextAction {
  type: typeof SEARCH_BAR_UPDATE_TEXT
  text: string
}

interface ExpandAction {
  type: typeof SEARCH_BAR_EXPAND
}

interface CollapseAction {
  type: typeof SEARCH_BAR_COLLAPSE
}

interface SetShowSuggestionsAction {
  type: typeof SEARCH_BAR_SET_SHOW_SUGGESTIONS
  value: boolean
}

interface UpdateTextSuggestionsAction {
  type: typeof SEARCH_BAR_UPDATE_TEXT_SUGGESTIONS
  suggestions: string[]
}

export type SearchBarActionTypes =
  | UpdateTextAction
  | ExpandAction
  | CollapseAction
  | SetShowSuggestionsAction
  | UpdateTextSuggestionsAction
