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
}

type Props = StateToProps & DispatchToProps & OwnProps

const EpisodeThumbnail: React.FC<Props> = ({
  podcast,
  episode,
  playEpisode,
}) => {
  const PlayIcon = iconMap['play']

  return (
    <div>
      <div
        className="relative md:w-28 w-24 md:h-28 h-24 mb-2 rounded-lg border cursor-pointer"
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
            'absolute left-0 top-0 md:w-28 w-24 md:h-28 h-24',
            'thumbnail-overlay',
          )}
        >
          <div
            className="flex items-center md:w-14 w-12 md:h-14 h-12 mx-auto rounded-full"
            style={{ background: 'rgba(255, 255, 255, 0.75)' }}
          >
            <PlayIcon className="md:w-10 w-8 md:h-10 h-8 ml-3 mx-auto fill-current text-gray-800" />
          </div>
        </div>
        {/* Duration */}
        <div className="absolute right-0 bottom-0 px-2 text-2xs font-medium tracking-wide text-gray-800 bg-gray-300 rounded-lg">
          {formatDuration(episode.duration)}
        </div>
      </div>

      <div
        className={classnames('relative w-full h-1 bg-gray-400 rounded-full', {
          hidden: episode.lastPlayedAt === '',
        })}
      >
        <div
          className="absolute h-1 top-0 left-0 bg-green-700 rounded-full"
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
