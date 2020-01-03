import { PodcastLink } from 'components/link'
import React from 'react'
import { Podcast, PodcastSearchResult } from 'types/app'
import { getImageUrl } from 'utils/dom'

export interface StateToProps {
  podcast: Podcast
  podcastSearchResult: PodcastSearchResult
}

export interface OwnProps {
  podcastId: string
  searchQuery: string
}

const ResultPodcastItem: React.FC<StateToProps & OwnProps> = ({
  podcast,
  podcastSearchResult,
}) => {
  return (
    <div className="flex mb-12">
      <div className="flex-none mr-1">
        <img
          className="md:w-24 w-16 md:h-24 w-16 object-contain rounded-lg border cursor-default"
          src={getImageUrl(podcast.urlParam)}
        />
      </div>

      <div className="md:pl-4 pl-1">
        <PodcastLink podcastUrlParam={podcast.urlParam}>
          <a
            className="line-clamp-2"
            dangerouslySetInnerHTML={{
              __html: podcastSearchResult.title || podcast.title,
            }}
          />
        </PodcastLink>

        <div
          className="line-clamp-1"
          dangerouslySetInnerHTML={{
            __html: podcastSearchResult.author || podcast.author,
          }}
        />
      </div>
    </div>
  )
}

export default ResultPodcastItem
