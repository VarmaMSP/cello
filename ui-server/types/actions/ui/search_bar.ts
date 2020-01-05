export const SEARCH_BAR_UPDATE_TEXT = 'search_bar/update_text'
export const SEARCH_BAR_UPDATE_TEXT_SUGGESTIONS =
  'search_bar/update_text_suggestions'

export interface UpdateTextAction {
  type: typeof SEARCH_BAR_UPDATE_TEXT
  text: string
}

export interface UpdateTextSuggestionsAction {
  type: typeof SEARCH_BAR_UPDATE_TEXT_SUGGESTIONS
  suggestions: string[]
}

export type SearchBarActionTypes =
  | UpdateTextAction
  | UpdateTextSuggestionsAction
