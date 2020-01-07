import { connect } from 'react-redux'
import { getText } from 'selectors/ui/search_bar'
import { getResultType, getSortBy } from 'selectors/ui/search_results_list'
import { AppState } from 'store'
import SearchResultsFilter, { StateToProps } from './search_results_filter'

function mapStateToProps(state: AppState): StateToProps {
  return {
    searchBarText: getText(state),
    resultType: getResultType(state),
    sortBy: getSortBy(state),
  }
}

export default connect<StateToProps, {}, {}, AppState>(mapStateToProps)(
  SearchResultsFilter,
)
