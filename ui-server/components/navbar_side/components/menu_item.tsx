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
      className={classnames(
        'flex items-center w-full h-full px-4 rounded-full my-2 cursor-pointer',
        {
          'text-gray-700': !active,
          'text-green-700': active,
          'bg-gray-200': active,
        },
      )}
    >
      <Icon className="w-4 h-4 mr-3 fill-current" />
      <h4 className="capitalize text-base leading-loose tracking-wide">
        {name}
      </h4>
    </a>
  )
}

export default MenuItem
