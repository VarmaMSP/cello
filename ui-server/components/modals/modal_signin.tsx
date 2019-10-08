import { iconMap } from 'components/icon'
import React from 'react'
import ButtonSocialSignin from './components/button_social_signin'
import ModalContainer from './components/modal_container'
import Overlay from './components/overlay'

export interface Props {
  closeModal: () => void
}

const ModalSignin: React.SFC<Props> = (props) => {
  const LogoIcon = iconMap['logo-lg']

  return (
    <Overlay background="rgba(255, 255, 255, 0.95)">
      <ModalContainer handleClose={props.closeModal} closeUponClicking="CROSS">
        <div className="h-full pt-8">
          <LogoIcon className="mx-auto" />
          <div className="text-center mt-3 mb-10">
            Subscribe to podcasts, create playlists and much more
          </div>
          <ButtonSocialSignin
            icon="google-color"
            text="Sign in with Google"
            onClick={() => {
              window.location.href = `/api/signin/google`
            }}
          />
          <ButtonSocialSignin
            icon="facebook-color"
            text="Sign in with Faceebook"
            onClick={() => {
              window.location.href = `/api/signin/facebook`
            }}
          />
          <ButtonSocialSignin
            icon="twitter-color"
            text="Sign in with Twitter"
            onClick={() => {
              window.location.href = `/api/signin/twitter`
            }}
          />
        </div>
      </ModalContainer>
    </Overlay>
  )
}

export default ModalSignin
