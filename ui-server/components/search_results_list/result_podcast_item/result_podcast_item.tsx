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
            className="md:text-base text-sm font-medium tracking-wide line-clamp-2"
            dangerouslySetInnerHTML={{
              __html: podcastSearchResult.title || podcast.title,
            }}
          />
        </PodcastLink>

        <div
          className="text-sm text-grey-800 hover:text-black tracking-wide line-clamp-1"
          style={{ margin: '3px 0px' }}
          dangerouslySetInnerHTML={{
            __html: podcastSearchResult.author || podcast.author,
          }}
        />

        <div
          className="mt-1 text-xs text-gray-700 leading-snug tracking-wider line-clamp-2"
          style={{ hyphens: 'auto' }}
          dangerouslySetInnerHTML={{
            __html: podcastSearchResult.description || podcast.summary,
          }}
        />
      </div>
    </div>
  )
}

export default ResultPodcastItem
