import classNames from 'classnames'
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
    <div className="flex items-center">
      <div className="w-36 text-xs text-gray-700">
        {formatEpisodePubDate(episode.pubDate)}
        <span className="mx-2 font-extrabold">&middot;</span>
        {formatEpisodeDuration(episode.duration)}
      </div>
      <div
        className={classNames(
          'relative md:h-0.8 h-0.6 bg-gray-400 rounded-full',
          { 'lg:w-1/3 w-2/5': !!playback, 'w-0': !!!playback },
        )}
      >
        <div
          className="absolute top-0 left-0 md:h-0.8 h-0.6 bg-red-500 rounded-full"
          style={{
            transition: 'ease-in 0.4s',
            width: playback
              ? `${(playback.currentTime / episode.duration) * 100}%`
              : `0`,
          }}
        />
      </div>
    </div>
  )
}

export default EpisodeMeta
