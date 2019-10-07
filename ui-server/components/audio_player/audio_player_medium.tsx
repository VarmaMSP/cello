import ButtonWithIcon from 'components/button_with_icon'
import React from 'react'
import { AudioState, Episode, Podcast } from 'types/app'
import ActionButton from './components/action_button'
import NavbarBottom from './components/navbar_bottom'
import SeekBar from './components/seek_bar'

interface Props {
  episode: Episode
  podcast: Podcast
  audioState: AudioState
  duration: number
  currentTime: number
  handleSeek: (t: number) => void
  handleFastForward: (t: number) => void
  handleActionButtonPress: () => void
}

const AudioPlayerMedium: React.SFC<Props> = (props) => {
  const {
    podcast,
    episode,
    audioState,
    duration,
    currentTime,
    handleSeek,
    handleActionButtonPress,
    handleFastForward,
  } = props

  const player = episode ? (
    <div className="flex items-center justify-center w-full h-20">
      <div className="flex-none flex items-center h-full">
        <ButtonWithIcon
          className="w-8"
          icon="fast-rewind"
          onClick={() => handleFastForward(-10)}
        />
        <ActionButton
          audioState={audioState}
          handleActionButtonPress={handleActionButtonPress}
        />
        <ButtonWithIcon
          className="w-8"
          icon="fast-forward"
          onClick={() => handleFastForward(10)}
        />
      </div>
      <div className="flex-auto flex flex-col justify-center ml-6 text-center truncate">
        <div className="-mb-1">
          <div className="text-sm font-semibold text-gray-800 leading-tight tracking-wide">
            {episode.title}
          </div>
          <div className="text-sm text-gray-700 leading-loose tracking-wide truncate">
            {podcast.title}
          </div>
        </div>
        <div className="w-full">
          <SeekBar
            currentTime={currentTime}
            duration={duration}
            handleSeek={handleSeek}
          />
        </div>
      </div>
    </div>
  ) : (
    <></>
  )

  return (
    <div className="fixed bottom-0 left-0 flex flex-col justify-between bg-white w-full h-auto px-5 border">
      {player}
      <div className="mt-2">
        <NavbarBottom />
      </div>
    </div>
  )
}

export default AudioPlayerMedium
