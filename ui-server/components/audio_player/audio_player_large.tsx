import React from 'react'
import { AudioState } from './components/utils'
import ButtonWithIcon from '../button_with_icon'
import ActionButton from './components/action_button'
import ProgressBar from './components/progress_bar'

interface Props {
  podcast: string
  episode: string
  podcastId: string
  episodeId: string
  albumArt: string
  audioState: AudioState
  duration: number
  currentTime: number
  handleSeek: (t: number) => void
  handleFastForward: (t: number) => void
  handleActionButtonPress: () => void
}

const AudioPlayerLarge: React.SFC<Props> = (props) => {
  const {
    podcast,
    episode,
    audioState,
    currentTime,
    duration,
    handleSeek,
    handleActionButtonPress,
    handleFastForward,
  } = props

  return (
    <div className="fixed bottom-0 left-0 flex items-center justify-around w-full h-24 pl-56 bg-white border">
      <div className="flex-none flex items-center h-full mx-4">
        <ButtonWithIcon
          className="w-10"
          icon="fast-rewind"
          onClick={() => handleFastForward(-10)}
        />
        <ActionButton
          audioState={audioState}
          handleActionButtonPress={handleActionButtonPress}
        />
        <ButtonWithIcon
          className="w-10"
          icon="fast-forward"
          onClick={() => handleFastForward(10)}
        />
      </div>
      <div className="flex-auto flex flex-col justify-center mx-4 text-center">
        <div className="-mb-1">
          <div className="text-ms font-semibold text-gray-800 leading-tight tracking-wide truncate">
            {episode}
          </div>
          <div className="text-sm font-semibold text-gray-700 leading-loose tracking-tigh truncate">
            {podcast}
          </div>
        </div>
        <div className="w-full">
          <ProgressBar
            currentTime={currentTime}
            duration={duration}
            handleSeek={handleSeek}
          />
        </div>
      </div>
    </div>
  )
}

export default AudioPlayerLarge
