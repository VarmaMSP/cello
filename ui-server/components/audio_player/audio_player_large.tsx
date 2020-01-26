import ButtonWithIcon from 'components/button_with_icon'
import React from 'react'
import { AudioState, Episode, Podcast } from 'types/app'
import { getImageUrl } from 'utils/dom'
import { formatPlaybackRate, formatVolume } from 'utils/format'
import ActionButton from './components/action_button'
import RangeControl from './components/range_control'
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
    volume,
    playbackRate,
    handleSeek,
    handleFastForward,
    handleVolumeChange,
    handlePlaybackRateChange,
    handleActionButtonPress,
  } = props

  if (!episode) {
    return <></>
  }

  return (
    <div
      className="flex align-bottom md:w-1/2"
      style={{
        position: 'fixed',
        bottom: '0',
        right: '0',
        width: '50%',
        background: 'rgba(255, 255, 255, 0.55)',
      }}
    >
      <div
        className="flex items-center justify-around w-full h-22 border border-gray-200 rounded-lg"
        style={{
          marginBottom: '8px',
          marginRight: '8px',
          background: 'rgba(255, 255, 255)',
          boxShadow: '0 3px 6px rgba(0,0,0,0.23), 0 3px 6px rgba(0,0,0,0.16)',
        }}
      >
        <img
          className="flex-none w-22 h-22 border-r rounded-lg"
          src={getImageUrl(podcast.urlParam)}
        />

        <div className="flex-1 flex flex-col justify-around h-full px-4 py-1">
          <div className="flex justify-between">
            {/* Title */}
            <div className="flex-auto pr-4">
              <div className="text-sm text-gray-900 font-medium tracking-wide line-clamp-1">
                {episode.title}
              </div>
              <div className="text-xs text-gray-800 tracking-wide line-clamp-1">
                {podcast.title}
              </div>
            </div>

            {/* Controls */}
            <div className="flex-none flex items-center justify-around w-30">
              <ButtonWithIcon
                className="w-8 h-8 px-2 py-2 hover:bg-gray-200 rounded-full"
                icon="fast-rewind"
                onClick={() => handleFastForward(-10)}
              />
              <ActionButton
                audioState={audioState}
                handleActionButtonPress={handleActionButtonPress}
              />
              <ButtonWithIcon
                className="w-8 h-8 px-2 py-2 hover:bg-gray-200 rounded-full"
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

        <div className="hidden">
          <RangeControl
            icon="volume"
            value={volume}
            min={0}
            max={1}
            step={0.1}
            onChange={handleVolumeChange}
            formatValue={formatVolume}
          />
          <div className="my-4" />
          <RangeControl
            icon="walk"
            value={playbackRate}
            min={0.25}
            max={2.0}
            step={0.1}
            onChange={handlePlaybackRateChange}
            formatValue={formatPlaybackRate}
          />
        </div>
      </div>
    </div>
  )
}

export default AudioPlayerLarge
