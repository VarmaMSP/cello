import classNames from 'classnames'
import format from 'date-fns/format'
import parseISO from 'date-fns/parseISO'
import React from 'react'
import { Episode } from 'types/app'
import { formatEpisodeDuration } from 'utils/format'

export interface StateToProps {
  episode: Episode
}

export interface OwnProps {
  episodeId: string
  displayPubDate?: boolean
  displayDuration?: boolean
}

interface Props extends StateToProps, OwnProps {}

const EpisodeMeta: React.SFC<Props> = ({
  episode,
  displayPubDate = true,
  displayDuration = true,
}) => {
  let pubDate: string | undefined
  try {
    pubDate = format(parseISO(`${episode.pubDate} +0000`), 'PP')
  } catch (err) {}

  let duration: string | undefined
  if (episode.duration > 0) {
    duration = formatEpisodeDuration(episode.duration)
  }

  return (
    <div className="flex items-center">
      <div className="flex-none w-42 mr-2 text-2xs leading-relaxed text-gray-800">
        {displayPubDate && pubDate}
        {displayPubDate && displayDuration && pubDate && duration && (
          <span className="mx-2 font-extrabold">&middot;</span>
        )}
        {displayDuration && duration}
      </div>
      <div
        className={classNames(
          'relative flex-initial lg:w-1/3 w-2/5 bg-gray-400 rounded-full',
          { hidden: episode.lastPlayedAt === '' },
        )}
        style={{ height: '0.20rem' }}
      >
        <div
          className="absolute top-0 left-0 md:h-0.8 h-0.6 bg-red-500 rounded-full"
          style={{
            transition: 'ease-in 0.4s',
            height: '0.20rem',
            width: `${episode.progress}%`,
          }}
        />
      </div>
    </div>
  )
}

export default EpisodeMeta
