import { getSearchResults } from '../../selectors/entities/search'
import SearchResults, { StateToProps, OwnProps } from './search_results'
import { AppState } from '../../store'
import { connect } from 'react-redux'

function mapStateToProps(state: AppState): StateToProps {
  return {
    podcasts: getSearchResults()(state),
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  SearchResults,
)
