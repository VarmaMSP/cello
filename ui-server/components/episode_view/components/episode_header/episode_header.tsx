import { iconMap } from 'components/icon'
import { PodcastLink } from 'components/link'
import format from 'date-fns/format'
import parseISO from 'date-fns/parseISO'
import React from 'react'
import { Episode, Podcast } from 'types/app'
import { getImageUrl } from 'utils/dom'
import { formatEpisodeDuration } from 'utils/format'

export interface StateToProps {
  podcast: Podcast
}

export interface DispatchToProps {
  playEpisode: (stateTime: number) => void
  showAddToPlaylistModal: () => void
}

export interface OwnProps {
  episode: Episode
}

type Props = StateToProps & DispatchToProps & OwnProps

const EpisodeHeader: React.FC<Props> = ({
  episode,
  podcast,
  playEpisode,
  showAddToPlaylistModal,
}) => {
  let pubDate: string | undefined
  try {
    pubDate = format(parseISO(`${episode.pubDate} +0000`), 'PP')
  } catch (err) {}

  let duration: string | undefined
  if (episode.duration > 0) {
    duration = formatEpisodeDuration(episode.duration)
  }

  const PlayIcon = iconMap['play']
  const AddToPlaylistIcon = iconMap['playlist-add']

  return (
    <div className="flex">
      <img
        className="lg:h-36 h-24 lg:w-36 w-24 flex-none object-contain object-center rounded-lg border"
        src={getImageUrl(podcast.urlParam)}
      />

      <div className="flex flex-col flex-auto w-1/2 justify-between lg:px-5 px-3">
        <div className="w-full mb-3">
          <h2 className="text-lg text-gray-900 leading-tight line-clamp-2">
            {episode.title}
          </h2>
          <PodcastLink podcastUrlParam={podcast.urlParam}>
            <a className="md:text-base text-sm text-gray-800 hover:text-gray-900 leading-loose truncate">
              {podcast.title}
            </a>
          </PodcastLink>

          <div className="text-xs leading-relaxed tracking-wide text-gray-800">
            {`Published on ${pubDate}`}
            <span className="mx-2 font-extrabold">&middot;</span>
            {duration}
          </div>
        </div>

        <div className="flex">
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
        </div>
      </div>
    </div>
  )
}

export default EpisodeHeader
