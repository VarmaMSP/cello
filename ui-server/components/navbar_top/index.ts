import { connect } from 'react-redux'
import { Dispatch, bindActionCreators } from 'redux'
import { getIsUserSignedIn } from 'selectors/session'
import { getText } from 'selectors/ui/search_bar'
import { AppState } from 'store'
import * as T from 'types/actions'
import NavbarTop, { DispatchToProps, StateToProps } from './navbar_top'
import { loadResultsPage } from 'actions/results'

function mapStateToProps(state: AppState): StateToProps {
  return {
    userSignedIn: getIsUserSignedIn(state),
    searchText: getText(state),
  }
}

function mapDispatchToProps(dispatch: Dispatch<T.AppActions>): DispatchToProps {
  return {
    searchTextChange: (text: string) =>
      dispatch({ type: T.SEARCH_BAR_UPDATE_TEXT, text }),
    loadResultsPage: bindActionCreators(loadResultsPage, dispatch)
  }
}

export default connect<StateToProps, DispatchToProps, {}, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(NavbarTop)
