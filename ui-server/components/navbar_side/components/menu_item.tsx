import classnames from 'classnames'
import { Icon, iconMap } from 'components/icon'
import React from 'react'

interface Props {
  icon: Icon
  name: string
  active?: boolean
}

const MenuItem: React.SFC<Props> = ({ icon, name, active }) => {
  const Icon = iconMap[icon]

  return (
    <a
      href="#"
      className={classnames(
        'flex items-center w-full px-4 rounded-full',
        { 'text-gray-700': !active, 'text-green-600': active },
        { 'hover:bg-gray-200': !active },
        { 'cursor-pointer': !active, 'cursor-default': active },
      )}
    >
      <Icon className="w-5 h-5 mr-3 fill-current" />
      <h4 className="capitalize text-base leading-loose tracking-wide">
        {name}
      </h4>
    </a>
  )
}

export default MenuItem
