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
  removeEpisode: () => void
}

export interface OwnProps {
  position: number
  playlistId: string
  episodeId: string
}

type Props = StateToProps & DispatchToProps & OwnProps

const PlaylistEpisodeListItem: React.FC<Props> = ({
  episode,
  podcast,
  position,
  playEpisode,
  removeEpisode,
}) => {
  const PlayIcon = iconMap['play-outline']
  const DeleteIcon = iconMap['trash']

  return (
    <div className="flex items-center py-2">
      <div className="w-6 flex-none md:ml-2 md:mr-2 text-sm text-gray-600">
        {position}
      </div>
      <img
        className="md:w-14 w-12 md:h-14 h-12 md:mr-4 mr-2 flex-none object-contain rounded border"
        src={getImageUrl(podcast.urlParam)}
      />
      <div className="flex-auto">
        <EpisodeLink episodeUrlParam={episode.urlParam}>
          <a className="inline mb-2 md:text-base text-sm text-gray-900 tracking-wide leading-tight line-clamp-1">
            {episode.title}
          </a>
        </EpisodeLink>

        <PodcastLink podcastUrlParam={podcast.urlParam}>
          <a className="md:text-sm text-xs text-gray-700 tracking-wide line-clamp-1">
            {podcast.title}
          </a>
        </PodcastLink>
      </div>
      <button onClick={playEpisode} className="flex-none md:ml-4 ml-2">
        <PlayIcon className="md:w-6 w-5 h-auto fill-current text-gray-700" />
      </button>
      <button onClick={removeEpisode} className="flex-none md:ml-4 ml-2">
        <DeleteIcon className="w-4 h-auto fill-current text-gray-600" />
      </button>
    </div>
  )
}

export default PlaylistEpisodeListItem
