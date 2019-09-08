import { connect } from 'react-redux'
import { getSearchResults } from 'selectors/entities/search'
import { AppState } from 'store'
import SearchResults, { OwnProps, StateToProps } from './search_results'

function mapStateToProps(state: AppState): StateToProps {
  return {
    podcasts: getSearchResults()(state),
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  SearchResults,
)
