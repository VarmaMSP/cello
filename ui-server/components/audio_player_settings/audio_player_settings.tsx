import RangeControl from 'components/audio_player/components/range_control'
import React from 'react'
import { formatPlaybackRate, formatVolume } from 'utils/format'

export interface StateToProps {
  volume: number
  playbackRate: number
}

export interface DispatchToProps {
  changeVolume: (v: number) => void
  changePlaybackRate: (r: number) => void
}

const AudioPlayerSettings: React.FC<StateToProps & DispatchToProps> = ({
  volume,
  playbackRate,
  changeVolume,
  changePlaybackRate,
}) => {
  return (
    <div className="px-4 py-6 bg-white border border-gray-400 shadow-md rounded-lg">
      <RangeControl
        icon="volume"
        value={volume}
        min={0}
        max={1}
        step={0.1}
        onChange={changeVolume}
        formatValue={formatVolume}
      />
      <div className="my-4" />
      <RangeControl
        icon="walk"
        value={playbackRate}
        min={0.25}
        max={2.0}
        step={0.1}
        onChange={changePlaybackRate}
        formatValue={formatPlaybackRate}
      />
    </div>
  )
}

export default AudioPlayerSettings
