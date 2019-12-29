import { iconMap } from 'components/icon'
import ModalContainer from 'components/modal/modal_container'
import Overlay from 'components/modal/overlay'
import React from 'react'
import SocialSignIn from './social_sign_in'

const SignInModal: React.FC<{}> = () => {
  const LogoIcon = iconMap['logo-lg']

  return (
    <Overlay background="rgba(0, 0, 0, 0.75)">
      <ModalContainer>
        <div className="h-full pt-8">
          <LogoIcon className="mx-auto" />
          <div className="text-center mt-3 mb-10">
            Subscribe to podcasts, create playlists and much more
          </div>
          <SocialSignIn
            icon="google-color"
            text="Sign in with Google"
            onClick={() => {
              window.location.href = `/api/signin/google`
            }}
          />
          <SocialSignIn
            icon="facebook-color"
            text="Sign in with Facebook"
            onClick={() => {
              window.location.href = `/api/signin/facebook`
            }}
          />
          <SocialSignIn
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

export default SignInModal
