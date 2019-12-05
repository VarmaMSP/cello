import EpisodeMeta from 'components/episode_meta'
import { iconMap } from 'components/icon'
import { EpisodeLink, PodcastLink } from 'components/link'
import React from 'react'
import striptags from 'striptags'
import { Episode, Podcast } from 'types/app'
import { getImageUrl } from 'utils/dom'

export interface StateToProps {
  episode: Episode
  podcast: Podcast
}

export interface DispatchToProps {
  playEpisode: (stateTime: number) => void
  showAddToPlaylistModal: () => void
}

export interface OwnProps {
  episodeId: string
}

const EpisodeListItem: React.FC<StateToProps & DispatchToProps & OwnProps> = ({
  episode,
  podcast,
  playEpisode,
  showAddToPlaylistModal,
}) => {
  const PlayIcon = iconMap['play']
  const AddToPlaylistIcon = iconMap['playlist-add']
  const ShareIcon = iconMap['share']

  return (
    <div className="flex mb-12">
      <img
        className="w-24 h-24 mr-2 flex-none object-contain rounded-lg border cursor-default"
        src={getImageUrl(episode.podcastId)}
      />
      <div className="pl-3">
        <EpisodeLink episodeId={episode.id}>
          <a className="md:text-base text-sm line-clamp-2">{episode.title}</a>
        </EpisodeLink>
        <PodcastLink podcastId={podcast.id}>
          <a className="text-sm text-grey-800 hover:text-black my-1">
            {podcast.title}
          </a>
        </PodcastLink>

        <EpisodeMeta episodeId={episode.id} />
        <EpisodeLink episodeId={episode.id}>
          <a
            className="mt-1 text-xs text-gray-700 leading-snug tracking-wide line-clamp-2"
            style={{ hyphens: 'auto' }}
          >
            {striptags(episode.description)}
          </a>
        </EpisodeLink>

        <div className="flex mt-4">
          <button
            className="flex items-center mr-4 px-3 py-1 text-2xs text-center text-purple-900 bg-gray-300 border hover:border-2 rounded-lg"
            onClick={() =>
              playEpisode((episode.progress * episode.duration) / 100)
            }
          >
            <PlayIcon className="fill-current w-4 h-auto" />
            <span className="ml-2 font-medium">PLAY</span>
          </button>
          <button
            className="flex items-center mr-4 px-3 py-1 text-2xs text-center text-gray-700 bg-gray-200 border rounded-lg"
            onClick={() => showAddToPlaylistModal()}
          >
            <AddToPlaylistIcon className="fill-current w-4 h-auto" />
            <span className="ml-2 font-medium">ADD</span>
          </button>
          <button className="flex items-center mr-4 px-3 py-1 text-2xs text-center text-gray-700 bg-gray-200 border rounded-lg">
            <ShareIcon className="fill-current w-3 h-auto" />
            <span className="ml-2 font-medium">SHARE</span>
          </button>
        </div>
      </div>
    </div>
  )
}

export default EpisodeListItem
