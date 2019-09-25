import { Dispatch } from 'react'
import { connect } from 'react-redux'
import { AppState } from 'store'
import { AppActions, CLOSE_SIGN_IN_MODAL } from 'types/actions'
import SigninModal, { DispatchToProps, StateToProps } from './modal_signin'

function mapStateToProps(state: AppState): StateToProps {
  return {
    showSignInModal: state.ui.showSignInModal,
  }
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return { closeModal: () => dispatch({ type: CLOSE_SIGN_IN_MODAL }) }
}

export default connect<StateToProps, DispatchToProps, {}, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(SigninModal)
