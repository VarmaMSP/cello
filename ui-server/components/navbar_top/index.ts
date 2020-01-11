import { loadResultsPage } from 'actions/results'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getIsUserSignedIn } from 'selectors/session'
import { getText } from 'selectors/ui/search_bar'
import { getResultType, getSortBy } from 'selectors/ui/search_results_list'
import { AppState } from 'store'
import * as T from 'types/actions'
import NavbarTop, { DispatchToProps, StateToProps } from './navbar_top'

function mapStateToProps(state: AppState): StateToProps {
  return {
    userSignedIn: getIsUserSignedIn(state),
    searchText: getText(state),
    resultType: getResultType(state),
    sortBy: getSortBy(state),
  }
}

function mapDispatchToProps(dispatch: Dispatch<T.AppActions>): DispatchToProps {
  return {
    searchTextChange: (text: string) =>
      dispatch({ type: T.SEARCH_BAR_UPDATE_TEXT, text }),
    loadResultsPage: bindActionCreators(loadResultsPage, dispatch),
  }
}

export default connect<StateToProps, DispatchToProps, {}, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(NavbarTop)
