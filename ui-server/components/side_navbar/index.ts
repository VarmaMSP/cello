import { connect } from 'react-redux'
import { Dispatch } from 'redux'
import { getIsUserSignedIn } from 'selectors/session'
import { getViewportSize } from 'selectors/window'
import { AppState } from 'store'
import { AppActions, MODAL_MANAGER_SHOW_SIGN_IN_MODAL } from 'types/actions'
import Navbar, { DispatchToProps, StateToProps } from './side_navbar'

function mapStateToProps(state: AppState): StateToProps {
  return {
    userSignedIn: getIsUserSignedIn(state),
    viewportSize: getViewportSize(state),
  }
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    showSigninModal: () => dispatch({ type: MODAL_MANAGER_SHOW_SIGN_IN_MODAL }),
  }
}

export default connect<StateToProps, DispatchToProps, {}, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(Navbar)
