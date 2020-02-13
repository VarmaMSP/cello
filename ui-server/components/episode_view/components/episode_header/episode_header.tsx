import EpisodeThumbnail from 'components/episode_thumbnail'
import { PodcastLink } from 'components/link'
import format from 'date-fns/format'
import parseISO from 'date-fns/parseISO'
import { Episode, Podcast } from 'models'
import React from 'react'

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

  return (
    <div className="flex">
      <EpisodeThumbnail episodeId={episode.id} large showIcon />
      <div className="flex flex-col flex-auto w-1/2 justify-between lg:px-5 px-3">
        <div className="w-full mb-3">
          <h2 className="md:text-xl text-lg text-gray-900 font-medium leading-snug line-clamp-2">
            {episode.title}
          </h2>
          <PodcastLink podcastUrlParam={podcast.urlParam}>
            <a className="block mb-2 md:text-base text-sm leading-relaxed line-clamp-1">
              {podcast.title}
            </a>
          </PodcastLink>

          <div className="text-2xs">{`Published on ${pubDate}`}</div>
        </div>

        <div className="flex align-bottom">
          <button
            className="w-42 py-1 text-xs text-center text-white font-semibold tracking-wide bg-purple-600 rounded focus:outline-none focus:shadow-outline"
            onClick={() => showAddToPlaylistModal()}
          >
            <span className="ml-2 font-medium">ADD TO PLAYLIST</span>
          </button>
        </div>
      </div>
    </div>
  )
}

export default EpisodeHeader
