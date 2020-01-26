import EpisodeThumbnail from 'components/episode_thumbnail'
import { iconMap } from 'components/icon'
import { PodcastLink } from 'components/link'
import format from 'date-fns/format'
import parseISO from 'date-fns/parseISO'
import React from 'react'
import { Episode, Podcast } from 'types/app'

export interface StateToProps {
  podcast: Podcast
}

export interface DispatchToProps {
  showAddToPlaylistModal: () => void
}

export interface OwnProps {
  episode: Episode
}

type Props = StateToProps & DispatchToProps & OwnProps

const EpisodeHeader: React.FC<Props> = ({
  episode,
  podcast,
  showAddToPlaylistModal,
}) => {
  let pubDate: string | undefined
  try {
    pubDate = format(parseISO(`${episode.pubDate} +0000`), 'PP')
  } catch (err) {}

  const AddToPlaylistIcon = iconMap['playlist-add']

  return (
    <div className="flex">
      <EpisodeThumbnail episodeId={episode.id} large showIcon />
      <div className="flex flex-col flex-auto w-1/2 justify-between lg:px-5 px-3">
        <div className="w-full mb-3">
          <h2 className="md:text-xl text-lg text-gray-900 font-medium leading-snug line-clamp-2">
            {episode.title}
          </h2>
          <PodcastLink podcastUrlParam={podcast.urlParam}>
            <a className="block mb-2 text-sm text-gray-800 truncate">
              {podcast.title}
            </a>
          </PodcastLink>

          <div className="text-xs text-gray-700">
            {`Published on ${pubDate}`}
          </div>
        </div>

        <div className="flex align-bottom">
          <button
            className="flex items-center mr-4 px-3 py-1 text-2xs text-center text-gray-700 bg-gray-200 border rounded-lg focus:outline-none focus:shadow-outline"
            onClick={() => showAddToPlaylistModal()}
          >
            <AddToPlaylistIcon className="fill-current w-4 h-auto" />
            <span className="ml-2 font-medium">ADD</span>
          </button>
        </div>
      </div>
    </div>
  )
}

export default EpisodeHeader
