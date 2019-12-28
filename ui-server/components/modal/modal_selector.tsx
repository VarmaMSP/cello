import AddToPlaylistModal from 'components/add_to_playlist_modal'
import CreatePlaylistModal from 'components/create_playlist_modal'
import { connect } from 'react-redux'
import { Dispatch } from 'redux'
import { getActiveModal } from 'selectors/ui/modal_manager'
import { AppState } from 'store'
import { AppActions, MODAL_MANAGER_CLOSE_MODAL } from 'types/actions'
import { Modal } from 'types/app'
import SignInModal from '../signin_modal'

interface StateToProps {
  modalToShow: Modal
}

interface DispatchToProps {
  closeModal: () => void
}

const ModalSelector: React.SFC<StateToProps & DispatchToProps> = ({
  modalToShow,
  closeModal,
}) => {
  switch (modalToShow.type) {
    case 'SIGNIN_MODAL':
      return <SignInModal closeModal={closeModal} />

    case 'ADD_TO_PLAYLIST_MODAL':
      return (
        <AddToPlaylistModal
          closeModal={closeModal}
          episodeId={modalToShow.episodeId}
        />
      )

    case 'CREATE_PLAYLIST_MODAL':
      return (
        <CreatePlaylistModal
          closeModal={closeModal}
          episodeId={modalToShow.episodeId}
        />
      )

    default:
      return <></>
  }
}

function mapStateToProps(state: AppState): StateToProps {
  return {
    modalToShow: getActiveModal(state),
  }
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    closeModal: () => dispatch({ type: MODAL_MANAGER_CLOSE_MODAL }),
  }
}

export default connect<StateToProps, DispatchToProps, {}, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(ModalSelector)
