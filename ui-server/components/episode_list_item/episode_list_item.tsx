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
      <div className="flex-none mr-1">
        <img
          className="w-24 h-24 object-contain rounded-lg border cursor-default"
          src={getImageUrl(podcast.urlParam)}
        />
      </div>

      <div className="pl-3">
        <EpisodeLink episodeUrlParam={episode.urlParam}>
          <a className="block md:text-base text-sm font-medium tracking-wide line-clamp-2">
            {episode.title}
          </a>
        </EpisodeLink>
        <PodcastLink podcastUrlParam={podcast.urlParam}>
          <a
            className="block text-xs text-grey-800 hover:text-black tracking-wide line-clamp-1"
            style={{ margin: '3px 0px' }}
          >
            {podcast.title}
          </a>
        </PodcastLink>

        <EpisodeMeta episodeId={episode.id} />
        <EpisodeLink episodeUrlParam={episode.urlParam}>
          <a
            className="mt-1 text-xs text-gray-700 leading-snug tracking-wider line-clamp-2"
            style={{ hyphens: 'auto' }}
          >
            {striptags(episode.description)}
          </a>
        </EpisodeLink>

        <div className="flex mt-4">
          <button
            className="flex items-center mr-4 px-3 py-1 text-2xs text-center bg-indigo-500 text-white border hover:border-2 rounded-lg focus:outline-none focus:shadow-outline"
            onClick={() =>
              playEpisode((episode.progress * episode.duration) / 100)
            }
          >
            <PlayIcon className="fill-current w-4 h-auto" />
            <span className="ml-1 font-medium">PLAY</span>
          </button>
          <button
            className="flex items-center mr-4 px-3 py-1 text-2xs text-center text-gray-700 bg-gray-200 border rounded-lg focus:outline-none focus:shadow-outline"
            onClick={() => showAddToPlaylistModal()}
          >
            <AddToPlaylistIcon className="fill-current w-4 h-auto" />
            <span className="ml-2 font-medium">ADD</span>
          </button>
          <button className="flex items-center mr-4 px-3 py-1 text-2xs text-center text-gray-700 bg-gray-200 border rounded-lg focus:outline-none focus:shadow-outline">
            <ShareIcon className="fill-current w-3 h-auto" />
            <span className="ml-2 font-medium">SHARE</span>
          </button>
        </div>
      </div>
    </div>
  )
}

export default EpisodeListItem
