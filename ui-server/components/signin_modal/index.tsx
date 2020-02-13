import ModalContainer from 'components/modal/modal_container'
import Overlay from 'components/modal/overlay'
import React from 'react'
import SocialSignIn from './social_sign_in'

const SignInModal: React.FC<{}> = () => {
  return (
    <Overlay background="rgba(0, 0, 0, 0.75)">
      <ModalContainer className="modal">
        <h1 className="text-5xl text-center text-purple-700 font-semibold tracking-wide">
          {'Phenopod'}
        </h1>
        <h2 className="max-w-sm mx-auto text-black text-center tracking-wide mt-4">
          {
            'subscribe to podcasts, create playlists, sync progress across multiple devices and much more'
          }
        </h2>
        <div className="h-full pt-6">
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
