import classnames from 'classnames'
import { iconMap } from 'components/icon'
import { Episode, Podcast } from 'types/models'
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
  small?: boolean
  large?: boolean
  showIcon?: boolean
}

type Props = StateToProps & DispatchToProps & OwnProps

const EpisodeThumbnail: React.FC<Props> = ({
  podcast,
  episode,
  playEpisode,
  small = false,
  large = false,
  showIcon = false,
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
            'md:h-36 md:w-36': large,
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
            'absolute left-0 top-0 w-22 h-22 flex items-center',
            {
              'md:w-25 md:h-25': !!small,
              'md:w-28 md:h-28': !!!small,
              'md:h-36 md:w-36': large,
            },
            {
              overlay: showIcon,
              'overlay-on-hover': !showIcon,
            },
          )}
        >
          <div
            className={classnames(
              'flex items-center w-10 h-10 mx-auto rounded-full',
              { 'md:w-14 md:h-14': !small },
            )}
            style={{ background: 'rgba(255, 255, 255, 0.75)' }}
          >
            <PlayIcon
              className={classnames('w-7 h-7 ml-2 fill-current text-gray-800', {
                'md:w-10 md:h-10 md:ml-3': !small,
              })}
            />
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
        className={classnames('relative w-full h-1 mt-2 bg-gray-400 rounded-full', {
          hidden: episode.lastPlayedAt === '',
        })}
      >
        <div
          className="absolute top-0 left-0 h-1 bg-red-700 rounded-full"
          style={{
            transition: 'ease-in 0.4s',
            width: `${episode.progress}%`,
          }}
        ></div>
      </div>
    </div>
  )
}

export default EpisodeThumbnail
