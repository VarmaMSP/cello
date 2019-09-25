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

  return <></>
}

export default Modals
