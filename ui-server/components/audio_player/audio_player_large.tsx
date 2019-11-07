import ButtonWithIcon from 'components/button_with_icon'
import Link from 'next/link'
import React from 'react'
import { AudioState, Episode, Podcast } from 'types/app'
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
    <div className="fixed bottom-0 left-0 flex items-center justify-around w-full h-24 pl-56 bg-white border">
      <div className="flex-none flex items-center h-full mx-4">
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
      <div className="flex-auto flex flex-col justify-center mx-4 text-center">
        <div className="-mb-1">
          <div className="text-base font-semibold text-gray-800 leading-wide tracking-wide truncate">
            {episode.title}
          </div>
          <Link
            href={{ pathname: '/podcasts', query: { podcastId: podcast.id, activeTab: 'episodes' } }}
            as={`/podcasts/${podcast.id}/episodes`}
            key={podcast.id}
          >
            <a className="block text-xs font-semibold text-gray-700 leading-loose tracking-tight truncate">
              {podcast.title}
            </a>
          </Link>
        </div>
        <div className="w-full">
          <SeekBar
            currentTime={currentTime}
            duration={duration}
            handleSeek={handleSeek}
          />
        </div>
      </div>
      <div className="mx-4">
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
  )
}

export default AudioPlayerLarge
