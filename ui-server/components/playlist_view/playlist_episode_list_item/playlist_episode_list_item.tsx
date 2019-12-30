import { iconMap } from 'components/icon'
import { EpisodeLink, PodcastLink } from 'components/link'
import React from 'react'
import { Episode, Podcast } from 'types/app'
import { getImageUrl } from 'utils/dom'

export interface StateToProps {
  episode: Episode
  podcast: Podcast
}

export interface DispatchToProps {
  playEpisode: () => void
}

export interface OwnProps {
  position: number
  episodeId: string
}

type Props = StateToProps & DispatchToProps & OwnProps

const PlaylistEpisodeListItem: React.FC<Props> = ({
  episode,
  podcast,
  position,
  playEpisode,
}) => {
  const PlayIcon = iconMap['play-outline']

  return (
    <div className="flex items-center py-2">
      <div className="w-6 md:ml-2 mr-2 text-sm text-gray-600">{position}</div>
      <img
        className="w-16 h-16 mr-4 flex-none object-contain rounded border"
        src={getImageUrl(podcast.urlParam)}
      />
      <div className="flex-auto">
        <EpisodeLink episodeUrlParam={episode.urlParam}>
          <a className="inline mb-2 text-gray-900 tracking-wide leading-tight line-clamp-1">
            {episode.title}
          </a>
        </EpisodeLink>

        <PodcastLink podcastUrlParam={podcast.urlParam}>
          <a className="text-gray-700 tracking-wide line-clamp-1">
            {podcast.title}
          </a>
        </PodcastLink>
      </div>
      <button onClick={playEpisode}>
        <PlayIcon className="w-6 h-auto fill-current text-gray-700" />
      </button>
    </div>
  )
}

export default PlaylistEpisodeListItem
