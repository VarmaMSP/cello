import { iconMap } from 'components/icon'
import React from 'react'
import Modal from './components/modal'
import SocialSignInButton from './components/social_sign_in_button'

export interface StateToProps {
  showSignInModal: boolean
}

export interface DispatchToProps {
  closeModal: () => void
}

interface Props extends StateToProps, DispatchToProps {}

const LoginModal: React.SFC<Props> = (props) => {
  const { showSignInModal, closeModal } = props
  const LogoIcon = iconMap['logo-lg']

  if (!showSignInModal) {
    return <></>
  }

  return (
    <Modal handleClose={closeModal}>
      <div className="h-full pt-8">
        <LogoIcon className="mx-auto" />
        <div className="text-center mt-3 mb-10">
          The Best web plodcast player
        </div>
        <SocialSignInButton icon="google-color" text="Sign in with Google" />
        <SocialSignInButton
          icon="facebook-color"
          text="Sign in with Faceebook"
        />
        <SocialSignInButton icon="twitter-color" text="Sign in with Twitter" />
      </div>
    </Modal>
  )
}

export default LoginModal
