import { loadResultsPage } from 'actions/results'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getQuery, getResultType, getSortBy } from 'selectors/ui/search_results_list'
import { AppState } from 'store'
import * as T from 'types/actions'
import SearchResultsFilter, { DispatchToProps, StateToProps } from './search_results_filter'

function mapStateToProps(state: AppState): StateToProps {
  return {
    query: getQuery(state),
    resultType: getResultType(state),
    sortBy: getSortBy(state),
  }
}

function mapDispatchToProps(dispatch: Dispatch<T.AppActions>): DispatchToProps {
  return {
    loadResultsPage: bindActionCreators(loadResultsPage, dispatch),
  }
}

export default connect<StateToProps, DispatchToProps, {}, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(SearchResultsFilter)
