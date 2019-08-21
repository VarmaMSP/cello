import React from 'react'
import ButtonWithIcon from '../../button_with_icon'
import { AudioState } from './utils';

interface Props {
  audioState: AudioState;
  handleActionButtonPress: () => void;
}

const ActionButton: React.SFC<Props> = ({audioState, handleActionButtonPress}) => {
  if (audioState === "PLAYING") {
    return <ButtonWithIcon
      className="w-16 h-16 px-3 py-3"
      icon="pause"
      onClick={handleActionButtonPress}
    />
  }

  if (audioState === "PAUSED") {
    return <ButtonWithIcon
      className="w-16 h-16 px-3 py- 3"
      icon="play"
      onClick={handleActionButtonPress}
    />
  }

  if (audioState === "LOADING") {
    return <div className="w-16 h-16 flex item-center">
      <div className="spinner-md mx-auto"/>
    </div>
  }
}

export default ActionButton