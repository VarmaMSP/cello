import { iconMap } from 'components/icon'
import Modal from 'components/modal'
import React from 'react'

const LoginModal: React.SFC<{}> = () => {
  const LogoIcon = iconMap['logo-lg']
  const GoogleIcon = iconMap['google-color']
  const FacebookIcon = iconMap['facebook-color']
  const TwitterIcon = iconMap['twitter-color']

  return (
    <Modal>
      <div className="h-full pt-8">
        <LogoIcon className="mx-auto" />
        <div className="text-center mt-3 mb-10">
          The Best web plodcast player
        </div>
        <button className="flex items-center justify-center md:w-3/5 w-full h-13 mx-auto my-4 border md:border-2 md:hover:border-gray-600 border-gray-400 rounded-lg">
          <GoogleIcon className="w-6" />
          <div className="text-lg text-gray-700 px-5">
            {'Sign in with Google'}
          </div>
        </button>
        <button className="flex items-center justify-center md:w-3/5 w-full h-13 mx-auto my-4 border md:border-2 md:hover:border-gray-600 border-gray-400 rounded-lg">
          <FacebookIcon className="w-6" />
          <div className="text-lg text-gray-700 px-5">
            {'Sign in with Facebook'}
          </div>
        </button>
        <button className="flex items-center justify-center md:w-3/5 w-full h-13 mx-auto my-4 border md:border-2 md:hover:border-gray-600 border-gray-400 rounded-lg">
          <TwitterIcon className="w-6" />
          <div className="text-lg text-gray-700 px-5">
            {'Sign in with Twitter'}
          </div>
        </button>
      </div>
    </Modal>
  )
}

export default LoginModal
