import { connect } from 'react-redux'
import { Dispatch } from 'redux'
import { AppState } from 'store'
import { AppActions, CLOSE_MODAL } from 'types/actions'
import Modal, { DispatchToProps, StateToProps } from './modals'

function mapStateToProps(state: AppState): StateToProps {
  return {
    modalToShow: state.ui.showModal,
  }
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    closeModal: () => dispatch({ type: CLOSE_MODAL }),
  }
}

export default connect<StateToProps, DispatchToProps, {}, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(Modal)
