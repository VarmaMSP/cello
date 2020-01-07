import { loadResultsPage } from 'actions/results'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getResultType, getSortBy } from 'selectors/ui/search_results_list'
import { AppState } from 'store'
import * as T from 'types/actions'
import { SearchResultType, SearchSortBy } from 'types/search'
import SearchResultsFilter, {
  DispatchToProps,
  StateToProps,
} from './search_results_filter'

function mapStateToProps(state: AppState): StateToProps {
  return {
    resultType: getResultType(state),
    sortBy: getSortBy(state),
  }
}

function mapDispatchToProps(dispatch: Dispatch<T.AppActions>): DispatchToProps {
  return {
    setResultType: (r: SearchResultType) =>
      dispatch({
        type: T.SEARCH_RESULTS_RESULT_TYPE,
        resultType: r,
      }),
    setSortBy: (s: SearchSortBy) =>
      dispatch({
        type: T.SEARCH_RESULTS_SORT_BY,
        sortBy: s,
      }),
    loadResultsPage: bindActionCreators(loadResultsPage, dispatch),
  }
}

export default connect<StateToProps, DispatchToProps, {}, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(SearchResultsFilter)
