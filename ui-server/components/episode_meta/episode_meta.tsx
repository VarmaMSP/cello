import React from 'react'
import { Episode, EpisodePlayback } from 'types/app'
import { formatEpisodeDuration, formatEpisodePubDate } from 'utils/format'

export interface StateToProps {
  playback?: EpisodePlayback
}

export interface OwnProps {
  episode: Episode
}

interface Props extends StateToProps, OwnProps {}

const EpisodeMeta: React.SFC<Props> = ({ episode, playback }) => {
  return (
    <div className="flex items-center justify-between">
      <span className="text-xs text-gray-700">
        {formatEpisodePubDate(episode.pubDate)}
        <span className="mx-2 font-extrabold">&middot;</span>
        {formatEpisodeDuration(episode.duration)}
      </span>
      {playback && (
        <div className="relative lg:w-1/3 w-2/5 h-1 bg-gray-400 rounded-full">
          <div
            className="absolute top-0 left-0 h-1 bg-red-500 rounded-full"
            style={{
              width: `${(playback.currentTime / episode.duration) * 100}%`,
            }}
          />
        </div>
      )}
    </div>
  )
}

export default EpisodeMeta
