import { PodcastLink } from 'components/link'
import React from 'react'
import { Episode, Podcast } from 'types/app'
import { getImageUrl } from 'utils/dom'

export interface StateToProps {
  podcast: Podcast
}

export interface OwnProps {
  episode: Episode
}

const EpisodeHeader: React.FC<StateToProps & OwnProps> = ({
  episode,
  podcast,
}) => {
  return (
    <div className="flex">
      <img
        className="lg:h-36 h-24 lg:w-36 w-24 flex-none object-contain object-center rounded-lg border"
        src={getImageUrl(podcast.urlParam)}
      />

      <div className="flex flex-col flex-auto w-1/2 justify-between lg:px-5 px-3">
        <div className="w-full">
          <h2 className="text-lg text-gray-900 leading-tight line-clamp-2">
            {episode.title}
          </h2>
          <PodcastLink podcastUrlParam={podcast.urlParam}>
            <a className="md:text-base text-sm text-gray-800 hover:text-gray-900 leading-loose truncate">
              {podcast.title}
            </a>
          </PodcastLink>
        </div>
      </div>
    </div>
  )
}

export default EpisodeHeader
