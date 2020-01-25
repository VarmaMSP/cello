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
        'flex items-center w-full h-full px-4 my-2 tracking-wide cursor-pointer',
        {
          'text-teal-800': !active,
          'text-red-600 bg-red-100 rounded-full': active,
        },
      )}
    >
      <Icon className="w-4 h-4 mr-3 fill-current" />
      <h4 className="capitalize text-lg font-medium leading-loose tracking-wide">
        {name}
      </h4>
    </a>
  )
}

export default MenuItem
