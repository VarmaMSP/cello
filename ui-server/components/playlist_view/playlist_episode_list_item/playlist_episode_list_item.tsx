import { EpisodeLink, PodcastLink } from 'components/link'
import React from 'react'
import { Episode, Podcast } from 'types/app'
import { getImageUrl } from 'utils/dom'

export interface StateToProps {
  episode: Episode
  podcast: Podcast
}

export interface OwnProps {
  position: number
  episodeId: string
}

const PlaylistEpisodeListItem: React.FC<StateToProps & OwnProps> = ({
  episode,
  podcast,
  position,
}) => {
  return (
    <div className="flex items-center py-2">
      <div className="w-6 md:ml-2 mr-2 text-sm text-gray-600">{position}</div>
      <img
        className="w-16 h-16 mr-4 flex-none object-contain rounded border"
        src={getImageUrl(podcast.urlParam)}
      />
      <div>
        <EpisodeLink episodeUrlParam={episode.urlParam}>
          <a className="block mb-2 text-gray-900 tracking-wide leading-tight line-clamp-1">
            {episode.title}
          </a>
        </EpisodeLink>
       
        <PodcastLink podcastUrlParam={podcast.urlParam}>
          <a className="block text-gray-700 tracking-wide line-clamp-1">
            {podcast.title}
          </a>
        </PodcastLink>
      </div>
    </div>
  )
}

export default PlaylistEpisodeListItem
