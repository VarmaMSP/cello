import AudioPlayerSettings from 'components/audio_player_settings'
import ButtonWithIcon from 'components/button_with_icon'
import { iconMap } from 'components/icon'
import { EpisodeLink } from 'components/link'
import usePopper from 'hooks/usePopper'
import React, { useState } from 'react'
import { Portal } from 'react-portal'
import { AudioState } from 'types/app'
import { Episode, Podcast } from 'types/models'
import { getImageUrl, stopEventPropagation } from 'utils/dom'
import ActionButton from './components/action_button'
import SeekBar from './components/seek_bar'

interface Props {
  episode: Episode
  podcast: Podcast
  audioState: AudioState
  duration: number
  currentTime: number
  volume: number
  playbackRate: number
  handleSeek: (t: number) => void
  handleFastForward: (t: number) => void
  handleVolumeChange: (v: number) => void
  handlePlaybackRateChange: (r: number) => void
  handleActionButtonPress: () => void
}

const AudioPlayerLarge: React.SFC<Props> = (props) => {
  const {
    podcast,
    episode,
    audioState,
    duration,
    currentTime,
    handleSeek,
    handleFastForward,
    handleActionButtonPress,
  } = props

  const [showPopper, setShowPopper] = useState<boolean>(false)
  const [reference, popper] = usePopper(
    {
      placement: 'top-end',
      modifiers: [
        {
          name: 'offset',
          options: {
            offset: [-6, 5],
          },
        },
      ],
      strategy: 'fixed',
    },
    () => setShowPopper(false),
  )

  if (!episode) {
    return <></>
  }

  const DotsIcon = iconMap['dots-horizontal']

  return (
    <div
      ref={reference.ref}
      className="fixed bottom-0 right-0 flex align-bottom md:w-1/2"
      style={{
        background: 'rgba(255, 255, 255, 0.55)',
      }}
    >
      <div
        className="flex items-center justify-between w-full h-22 border border-gray-200 bg-white rounded-lg"
        style={{
          marginBottom: '8px',
          marginRight: '8px',
          boxShadow: '0 3px 6px rgba(0,0,0,0.23), 0 3px 6px rgba(0,0,0,0.16)',
        }}
      >
        <img
          className="flex-none w-22 h-22 border-r rounded-lg"
          src={getImageUrl(podcast.urlParam)}
        />

        <div className="flex-1 flex flex-col justify-around h-full pl-4 pr-1 py-1">
          <div className="flex justify-between">
            {/* Title */}
            <EpisodeLink episodeUrlParam={episode.urlParam}>
              <a className="flex-auto pr-4">
                <div className="text-sm text-gray-900 font-medium tracking-wide line-clamp-1">
                  {episode.title}
                </div>
                <div className="text-xs text-gray-800 tracking-wide line-clamp-1">
                  {podcast.title}
                </div>
              </a>
            </EpisodeLink>

            {/* Controls */}
            <div className="flex-none flex items-center justify-around w-30">
              <ButtonWithIcon
                className="w-5 h-5"
                icon="fast-rewind"
                onClick={() => handleFastForward(-10)}
              />
              <ActionButton
                audioState={audioState}
                handleActionButtonPress={handleActionButtonPress}
              />
              <ButtonWithIcon
                className="w-5 h-5"
                icon="fast-forward"
                onClick={() => handleFastForward(10)}
              />
            </div>
          </div>

          <SeekBar
            currentTime={currentTime}
            duration={duration}
            handleSeek={handleSeek}
            compact
          />
        </div>

        <div className="flex-none h-full py-3">
          <button
            className="w-full px-2 py-3 focus:outline-none"
            onClick={() => setShowPopper(!showPopper)}
            onPointerDown={stopEventPropagation}
            onMouseDown={stopEventPropagation}
            onTouchStart={stopEventPropagation}
          >
            <DotsIcon className="w-4 h-4 mx-auto fill-current text-gray-800" />
          </button>
        </div>

        {showPopper && (
          <Portal>
            <div ref={popper.ref} style={popper.styles}>
              <AudioPlayerSettings />
            </div>
          </Portal>
        )}
      </div>
    </div>
  )
}

export default AudioPlayerLarge
