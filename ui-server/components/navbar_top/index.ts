import { connect } from 'react-redux'
import { Dispatch } from 'redux'
import { getIsUserSignedIn } from 'selectors/session'
import { getIsSearchBarCollapsed } from 'selectors/ui/search_bar'
import { getViewportSize } from 'selectors/window'
import { AppState } from 'store'
import * as T from 'types/actions'
import NavbarTop, { DispatchToProps, StateToProps } from './navbar_top'

function mapStateToProps(state: AppState): StateToProps {
  return {
    userSignedIn: getIsUserSignedIn(state),
    showSearchBar: !getIsSearchBarCollapsed(state),
    viewPortSize: getViewportSize(state),
  }
}

function mapDispatchToProps(dispatch: Dispatch<T.AppActions>): DispatchToProps {
  return {
    expandSearchBar: () => dispatch({ type: T.SEARCH_BAR_EXPAND }),
  }
}

export default connect<StateToProps, DispatchToProps, {}, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(NavbarTop)
