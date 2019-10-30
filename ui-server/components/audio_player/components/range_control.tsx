import ButtonWithIcon from 'components/button_with_icon'
import { Icon, iconMap } from 'components/icon'
import React from 'react'

interface Props {
  icon: Icon
  value: number
  max: number
  min: number
  step: number
  onChange: (val: number) => void
  formatValue: (val: number) => string
}

const epsilon = 1e-10

const RangeControl: React.SFC<Props> = (props) => {
  const { icon, value, max, min, step, onChange, formatValue } = props
  const Icon = iconMap[icon]

  return (
    <div className="flex items-center">
      <Icon className="w-4 h-auto mr-3 text-gray-700 fill-current" />
      <ButtonWithIcon
        icon="minus-outline"
        className="w-4 text-gray-600"
        onClick={() => value - step >= min - epsilon && onChange(value - step)}
      />
      <div className="w-12 text-xs text-gray-900 text-center">
        {formatValue(value)}
      </div>
      <ButtonWithIcon
        icon="add-outline"
        className="w-4 text-gray-600"
        onClick={() => value + step <= max + epsilon && onChange(value + step)}
      />
    </div>
  )
}

export default RangeControl
