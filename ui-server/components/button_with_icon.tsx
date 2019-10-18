import classnames from 'classnames'
import React from 'react'
import { Icon, iconMap } from './icon'

interface Props {
  icon: Icon
  className: string
  type?: 'button' | 'submit' | 'reset'
  onClick?: () => void
}

// A button containig only a svg icon.
//
// Icon
// - takes up the full width of the container while preserving its aspect ratio.
// - is centered vertically
// - fill is set to button's text-color
const ButtonWithIcon: React.SFC<Props> = (props) => {
  const { icon, className, type, onClick } = props
  const Icon = iconMap[icon]

  return (
    <button
      className={classnames(
        'flex items-center focus:outline-none cursor-pointer',
        className,
      )}
      type={type || 'button'}
      onClick={onClick}
    >
      <Icon className="fill-current w-full h-auto" />
    </button>
  )
}

export default ButtonWithIcon
