import { connect } from 'react-redux'
import { getText } from 'selectors/ui/search_bar'
import { AppState } from 'store'
import SearchResultsFilter, {
  OwnProps,
  StateToProps,
} from './search_results_filter'

function mapStateToProps(state: AppState): StateToProps {
  return {
    searchBarText: getText(state),
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  SearchResultsFilter,
)
