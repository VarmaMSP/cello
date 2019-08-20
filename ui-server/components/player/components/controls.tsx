import React from 'react'
import ButtonWithIcon from '../../button_with_icon'

interface Props {
  playing: boolean;
  handlePlayToggle: () => void;
  handleFastForward: (t: number) => void;
}

const Controls: React.SFC<Props> = ({playing, handlePlayToggle, handleFastForward}) => (
  <div className="flex justify-center items-center w-full py-1">
    <ButtonWithIcon
      className="w-8 h-8 mr-4"
      icon="fast-rewind"
      onClick={() => handleFastForward(-10)}
    />
    <ButtonWithIcon
      className="w-16 h-16 px-3 py-3 rounded-full shadow border"
      icon={playing ? "pause" : "play"}
      onClick={handlePlayToggle}
    />
    <ButtonWithIcon
      className="w-8 h-8 ml-4"
      icon="fast-forward"
      onClick={() => handleFastForward(10)}
    />
  </div>
)

export default Controls