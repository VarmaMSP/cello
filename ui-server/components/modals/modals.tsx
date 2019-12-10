import AddToPlaylistModal from 'components/add_to_playlist_modal'
import CreatePlaylistModal from 'components/create_playlist_modal'
import SignInModal from 'components/signin_modal'
import { Modal } from 'types/app'

export interface StateToProps {
  modalToShow: Modal
}

export interface DispatchToProps {
  closeModal: () => void
}

interface Props extends StateToProps, DispatchToProps {}

const Modals: React.SFC<Props> = ({ modalToShow, closeModal }) => {
  if (modalToShow.type === 'SIGNIN_MODAL') {
    return <SignInModal closeModal={closeModal} />
  }

  if (modalToShow.type === 'ADD_TO_PLAYLIST_MODAL') {
    return (
      <AddToPlaylistModal
        closeModal={closeModal}
        episodeId={modalToShow.episodeId}
      />
    )
  }

  if (modalToShow.type === 'CREATE_PLAYLIST_MODAL') {
    return (
      <CreatePlaylistModal
        closeModal={closeModal}
        episodeId={modalToShow.episodeId}
      />
    )
  }

  return <></>
}

export default Modals
