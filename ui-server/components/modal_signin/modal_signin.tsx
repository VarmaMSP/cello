import { iconMap } from 'components/icon'
import Modal from 'components/modal'
import React from 'react'
import ButtonSocialSignin from './components/button_social_signin'

export interface StateToProps {
  showSignInModal: boolean
}

export interface DispatchToProps {
  closeModal: () => void
}

interface Props extends StateToProps, DispatchToProps {}

const SigninModal: React.SFC<Props> = (props) => {
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
        <ButtonSocialSignin
          icon="google-color"
          text="Sign in with Google"
          onClick={() => {
            window.location.href = `${process.env.API_BASE_URL}/signin/google`
          }}
        />
        <ButtonSocialSignin
          icon="facebook-color"
          text="Sign in with Faceebook"
          onClick={() => {
            window.location.href = `${process.env.API_BASE_URL}/signin/facebook`
          }}
        />
        <ButtonSocialSignin
          icon="twitter-color"
          text="Sign in with Twitter"
          onClick={() => {
            window.location.href = `${process.env.API_BASE_URL}/signin/twitter`
          }}
        />
      </div>
    </Modal>
  )
}

export default SigninModal
