import { loadResultsPage } from 'actions/results'
import React, { createElement } from 'react'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getText } from 'selectors/ui/search_bar'
import { getResultType, getSortBy } from 'selectors/ui/search_results_list'
import { AppState } from 'store'
import * as T from 'types/actions'
import { SearchResultType, SearchSortBy } from 'types/search'

export interface SearchBarProps {
  searchText: string
  showSuggestions: boolean
  handleTextChange: (e: React.FormEvent<HTMLInputElement>) => void
  handleTextSubmit: (e?: React.FormEvent<HTMLFormElement>) => void
  collapseSearchBar: () => void
  setShowSuggestions: (e: boolean) => void
}

interface StateToProps {
  searchText: string
  resultType: SearchResultType
  sortBy: SearchSortBy
  showSuggestions: boolean
}

interface DispatchToProps {
  collapseSearchBar: () => void
  changeSearchText: (text: string) => void
  loadResultsPage: (
    query: string,
    resultType: SearchResultType,
    sortBy: SearchSortBy,
  ) => void
  setShowSuggestions: (e: boolean) => void
}

const withProps = (
  child: React.FC<SearchBarProps>,
): React.FC<StateToProps & DispatchToProps> => ({
  searchText,
  resultType,
  sortBy,
  collapseSearchBar,
  changeSearchText,
  loadResultsPage,
  showSuggestions,
  setShowSuggestions,
}) => {
  const handleTextChange = (e: React.FormEvent<HTMLInputElement>) => {
    changeSearchText(e.currentTarget.value)
  }

  const handleTextSubmit = (e?: React.FormEvent<HTMLFormElement>) => {
    !!e && e.preventDefault()
    collapseSearchBar()
    loadResultsPage(searchText, resultType, sortBy)
  }

  return createElement(child, {
    searchText,
    handleTextChange,
    handleTextSubmit,
    collapseSearchBar,
    showSuggestions,
    setShowSuggestions,
  })
}

const mapStateToProps = (state: AppState): StateToProps => ({
  searchText: getText(state),
  resultType: getResultType(state),
  sortBy: getSortBy(state),
  showSuggestions: state.ui.searchBar.showSuggestions,
})

const mapDispatchToProps = (
  dispatch: Dispatch<T.AppActions>,
): DispatchToProps => ({
  changeSearchText: (text: string) =>
    dispatch({ type: T.SEARCH_BAR_UPDATE_TEXT, text }),
  loadResultsPage: bindActionCreators(loadResultsPage, dispatch),
  collapseSearchBar: () => dispatch({ type: T.SEARCH_BAR_COLLAPSE }),
  setShowSuggestions: (e: boolean) =>
    dispatch({ type: T.SEARCH_BAR_SET_SHOW_SUGGESTIONS, value: e }),
})

export default (child: React.FC<SearchBarProps>) =>
  connect<StateToProps, DispatchToProps, {}, AppState>(
    mapStateToProps,
    mapDispatchToProps,
  )(withProps(child))
