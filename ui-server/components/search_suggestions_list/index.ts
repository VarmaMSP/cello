import { loadPodcastPage, loadResultsPage } from 'actions/results'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getSuggestions, getText } from 'selectors/ui/search_bar'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import { SearchSuggestion } from 'types/models'
import SearchSuggestionsList, {
  DispatchToProps,
  StateToProps,
} from './search_suggestions_list'

function mapStateToProps(state: AppState): StateToProps {
  const searchText = getText(state).trim()
  const currentSuggestion = <SearchSuggestion>{
    t: 'T',
    i: 'C',
    header: searchText,
    subHeader: '',
  }

  return {
    suggestions:
      searchText.length > 0
        ? [currentSuggestion, ...getSuggestions(state)]
        : getSuggestions(state),
  }
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
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
