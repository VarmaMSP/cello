import classnames from 'classnames'
import { iconMap } from 'components/icon'
import { Episode, Podcast } from 'types/app'
import { getImageUrl } from 'utils/dom'
import { formatDuration } from 'utils/format'

export interface StateToProps {
  episode: Episode
  podcast: Podcast
}

export interface DispatchToProps {
  playEpisode: (beginAt: number) => void
}

export interface OwnProps {
  episodeId: string
  small: boolean
  showIcon: boolean
}

type Props = StateToProps & DispatchToProps & OwnProps

const EpisodeThumbnail: React.FC<Props> = ({
  podcast,
  episode,
  playEpisode,
  small,
  showIcon,
}) => {
  const PlayIcon = iconMap['play']

  return (
    <div>
      <div
        className={classnames(
          'relative w-22 h-22 rounded-lg border cursor-pointer',
          {
            'md:w-25 md:h-25': !!small,
            'md:w-28 md:h-28': !!!small,
          },
        )}
        onClick={() => {
          episode.progress >= 95
            ? playEpisode(0)
            : playEpisode((episode.progress * episode.duration) / 100)
        }}
      >
        {/* Image */}
        <img
          className={classnames(
            'absolute left-0 top-0 right-0 bottom-0 w-full h-full',
            'object-contain rounded-lg cursor-default',
          )}
          src={getImageUrl(podcast.urlParam)}
        />

        {/* Icon */}
        <div
          className={classnames(
            'absolute left-0 top-0 w-22 h-22',
            {
              'md:w-25 md:h-25': !!small,
              'md:w-28 md:h-28': !!!small,
            },
            {
              overlay: showIcon,
              'overlay-on-hover': !showIcon,
            },
          )}
        >
          <div
            className="flex items-center md:w-14 md:h-14 w-10 h-10 mx-auto rounded-full"
            style={{ background: 'rgba(255, 255, 255, 0.75)' }}
          >
            <PlayIcon className="md:w-10 md:h-10 w-7 h-7 md:ml-3 ml-2 fill-current text-gray-800" />
          </div>
        </div>

        {/* Duration */}
        <div
          className="absolute right-0 bottom-0 px-1 text-2xs font-semibold text-gray-100 leading-tight rounded"
          style={{ background: 'rgba(1, 1, 1, 0.8)' }}
        >
          {formatDuration(episode.duration)}
        </div>
      </div>

      <div
        className={classnames('relative w-full mt-1 bg-gray-400 rounded-full', {
          hidden: episode.lastPlayedAt === '',
        })}
        style={{ height: '5px' }}
      >
        <div
          className="absolute top-0 left-0 bg-red-700 rounded-full"
          style={{
            transition: 'ease-in 0.4s',
            height: '5px',
            width: `${episode.progress}%`,
          }}
        ></div>
      </div>
    </div>
  )
}

export default EpisodeThumbnail
