import { Icon, iconMap } from 'components/icon'
import React from 'react'

interface Props {
  icon: Icon
  text: string
  onClick?: () => void
}

const SocialSignInButton: React.SFC<Props> = (props) => {
  const { icon, text, onClick } = props
  const SocialMediaIcon = iconMap[icon]

  return (
    <button
      onClick={onClick}
      className="flex items-center justify-center md:w-3/5 w-full h-13 mx-auto my-4 border md:border-2 md:hover:border-gray-600 border-gray-400 rounded-lg focus:outline-none focus:shadow-outline"
    >
      <SocialMediaIcon className="w-6" />
      <div className="text-lg text-gray-700 px-5">{text}</div>
    </button>
  )
}

export default SocialSignInButton
