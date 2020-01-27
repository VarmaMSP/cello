import classNames from 'classnames'
import ButtonWithIcon from 'components/button_with_icon'
import React from 'react'
import { AudioState } from 'types/app'

interface Props {
  theme?: string
  audioState: AudioState
  handleActionButtonPress: () => void
}

const ActionButton: React.SFC<Props> = ({
  theme,
  audioState,
  handleActionButtonPress,
}) => {
  if (audioState === 'PLAYING') {
    return (
      <ButtonWithIcon
        className={classNames('md:w-12 md:h-12 w-16 h-16 px-1 py-1', {
          'text-gray-300': theme === 'dark',
        })}
        icon="pause"
        onClick={handleActionButtonPress}
      />
    )
  }

  if (audioState === 'PAUSED' || audioState === 'ENDED') {
    return (
      <ButtonWithIcon
        className={classNames('md:w-12 md:h-12 w-16 h-16 px-1 py-1', {
          'text-gray-300': theme === 'dark',
        })}
        icon="play"
        onClick={handleActionButtonPress}
      />
    )
  }

  // return loader as default
  return (
    <div className="md:w-12 md:h-12 w-16 h-16 flex item-center cursor-wait">
      <div className="spinner mx-auto" />
    </div>
  )
}

export default ActionButton
