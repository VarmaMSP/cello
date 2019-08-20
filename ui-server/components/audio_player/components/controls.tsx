import React from 'react'
import ButtonWithIcon from '../../button_with_icon'
import { AudioState } from './utils';

interface Props {
  audioState: AudioState;
  handleFastForward: (t: number) => void;
  handleActionButtonPress: () => void;
}

const Controls: React.SFC<Props> = ({audioState, handleFastForward, handleActionButtonPress}) => {
  let actionButton: JSX.Element
  if (audioState === "PLAYING") {
    actionButton = <ButtonWithIcon
      className="w-16 h-16 px-3 py-3"
      icon="pause"
      onClick={handleActionButtonPress}
    />
  }
  if (audioState === "PAUSED") {
    actionButton =  <ButtonWithIcon
      className="w-16 h-16 px-3 py- 3"
      icon="play"
      onClick={handleActionButtonPress}
    />
  }
  if (audioState === "LOADING") {
    actionButton = <div className="w-16 h-16 flex item-center">
      <div className="spinner-md mx-auto"/>
    </div>
  }

  return <div className="flex justify-center items-center w-full py-1">
    <ButtonWithIcon
      className="w-8 h-8 mr-4"
      icon="fast-rewind"
      onClick={() => handleFastForward(-10)}
    />
    {actionButton}
    <ButtonWithIcon
      className="w-8 h-8 ml-4"
      icon="fast-forward"
      onClick={() => handleFastForward(10)}
    />
  </div>
}

export default Controls