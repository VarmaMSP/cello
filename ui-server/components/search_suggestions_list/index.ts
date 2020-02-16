import { loadPodcastPage, loadResultsPage } from 'actions/results'
import { SearchSuggestion } from 'models'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getCursor, getSuggestions, getText } from 'selectors/ui/search_bar'
import { AppState } from 'store'
import { AppActions, SEARCH_BAR_SET_CURSOR, SEARCH_BAR_UPDATE_TEXT } from 'types/actions'
import { uniqueId } from 'utils/utils'
import SearchSuggestionsList, { DispatchToProps, StateToProps } from './search_suggestions_list'

function mapStateToProps(state: AppState): StateToProps {
  return {
    cursor: getCursor(state),
    suggestions: [
      <SearchSuggestion>{
        id: uniqueId(),
        t: 'T',
        i: 'C',
        header: getText(state),
        subHeader: '',
      },
      ...getSuggestions(state),
    ],
  }
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    setCursor: (c: number) =>
      dispatch({
        type: SEARCH_BAR_SET_CURSOR,
        cursor: c,
      }),
    setSearchText: (t: string) =>
      dispatch({
        type: SEARCH_BAR_UPDATE_TEXT,
        text: t,
        skipSuggestions: true,
      }),
    loadResultsPage: (text: string) =>
      bindActionCreators(loadResultsPage, dispatch)(
        text,
        'episode',
        'relevance',
      ),
    loadPodcastPage: (podcastUrlParam: string) =>
      bindActionCreators(loadPodcastPage, dispatch)(podcastUrlParam),
  }
}

export default connect<StateToProps, DispatchToProps, {}, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(SearchSuggestionsList)
