import ModalAddToPlaylist from 'components/add_to_playlist_modal'
import CreatePlaylistModal from 'components/create_playlist_modal'
import { Modal } from 'types/app'
import ModalEpisode from './modal_episode'
import ModalSignin from './modal_signin'

export interface StateToProps {
  modalToShow: Modal
}

export interface DispatchToProps {
  closeModal: () => void
}

interface Props extends StateToProps, DispatchToProps {}

const Modals: React.SFC<Props> = ({ modalToShow, closeModal }) => {
  if (modalToShow.type === 'SIGNIN_MODAL') {
    return <ModalSignin closeModal={closeModal} />
  }

  if (modalToShow.type === 'EPISODE_MODAL') {
    const { episodeId } = modalToShow
    return <ModalEpisode closeModal={closeModal} episodeId={episodeId} />
  }

  if (modalToShow.type === 'ADD_TO_PLAYLIST_MODAL') {
    return <ModalAddToPlaylist episodeId={modalToShow.episodeId} />
  }

  if (modalToShow.type === 'CREATE_PLAYLIST_MODAL') {
    return <CreatePlaylistModal />
  }

  return <></>
}

export default Modals
