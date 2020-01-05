export interface SearchInput {
  query: string
  resultType: SearchResultType
  sortBy: SearchSortBy
}

export type SearchResultType = 'episode' | 'podcast'

export type SearchSortBy = 'relevance' | 'publish_date'
