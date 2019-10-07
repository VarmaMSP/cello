import ButtonWithIcon from 'components/button_with_icon'
import React from 'react'
import { AudioState } from 'types/app'

interface Props {
  audioState: AudioState
  handleActionButtonPress: () => void
}

const ActionButton: React.SFC<Props> = ({
  audioState,
  handleActionButtonPress,
}) => {
  if (audioState === 'PLAYING') {
    return (
      <ButtonWithIcon
        className="w-16 h-16 px-3 py-3"
        icon="pause"
        onClick={handleActionButtonPress}
      />
    )
  }

  if (audioState === 'PAUSED' || audioState === 'ENDED') {
    return (
      <ButtonWithIcon
        className="w-16 h-16 px-3 py- 3"
        icon="play"
        onClick={handleActionButtonPress}
      />
    )
  }

  // return loader as default
  return (
    <div className="w-16 h-16 flex item-center cursor-wait">
      <div className="spinner-md mx-auto" />
    </div>
  )
}

export default ActionButton
