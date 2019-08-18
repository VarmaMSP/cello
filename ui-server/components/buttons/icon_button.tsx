import React from 'react'
import { Icon, iconMap } from '../icons'
import classnames from 'classnames'

interface Props {
  icon: Icon,
  className: string
}

// A button with only a single svg icon inside it.
//
// Icon 
// - takes up the full width of the container while preserving its aspect ratio.
// - is centered vertically
// - fill is set to button's text-color 
const IconButton: React.SFC<Props> = ({icon, className}) => {
  let Icon = iconMap[icon]
  return <button className={classnames('flex center-items focus:outline-none', className)}>
    <Icon className='fill-current w-full h-auto'/>
  </button>
}

export default IconButton