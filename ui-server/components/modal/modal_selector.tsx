import AddToPlaylistModal from 'components/add_to_playlist_modal'
import CreatePlaylistModal from 'components/create_playlist_modal'
import { connect } from 'react-redux'
import { getActiveModal } from 'selectors/ui/modal_manager'
import { AppState } from 'store'
import { Modal } from 'types/app'
import SignInModal from '../signin_modal'

interface StateToProps {
  modalToShow: Modal
}

const ModalSelector: React.SFC<StateToProps> = ({ modalToShow }) => {
  switch (modalToShow.type) {
    case 'SIGNIN_MODAL':
      return <SignInModal />

    case 'ADD_TO_PLAYLIST_MODAL':
      return <AddToPlaylistModal episodeId={modalToShow.episodeId} />

    case 'CREATE_PLAYLIST_MODAL':
      return <CreatePlaylistModal episodeId={modalToShow.episodeId} />

    default:
      return <></>
  }
}

function mapStateToProps(state: AppState): StateToProps {
  return {
    modalToShow: getActiveModal(state),
  }
}

export default connect<StateToProps, {}, {}, AppState>(mapStateToProps)(
  ModalSelector,
)
