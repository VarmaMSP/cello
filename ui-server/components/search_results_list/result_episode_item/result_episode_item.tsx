import { EpisodeLink, PodcastLink } from 'components/link'
import React from 'react'
import { Episode, EpisodeSearchResult, Podcast } from 'types/app'
import { getImageUrl } from 'utils/dom'

export interface StateToProps {
  episode: Episode
  podcast: Podcast
  episodeSearchResult: EpisodeSearchResult
}

export interface OwnProps {
  episodeId: string
}

const ResultEpisodeItem: React.FC<StateToProps & OwnProps> = ({
  episode,
  podcast,
  episodeSearchResult,
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
        <EpisodeLink episodeUrlParam={episode.urlParam}>
          <a
            className="md:text-base text-sm font-medium tracking-wide line-clamp-2"
            dangerouslySetInnerHTML={{
              __html: episodeSearchResult.title || episode.title,
            }}
          />
        </EpisodeLink>

        <PodcastLink podcastUrlParam={podcast.urlParam}>
          <a
            className="text-sm text-grey-800 hover:text-black tracking-wide line-clamp-1"
            style={{ margin: '3px 0px' }}
          >
            {podcast.title}
          </a>
        </PodcastLink>

        <div
          className="mt-1 text-xs text-gray-700 leading-snug tracking-wider line-clamp-2"
          style={{ hyphens: 'auto' }}
          dangerouslySetInnerHTML={{
            __html: episodeSearchResult.description || episode.summary,
          }}
        />
      </div>
    </div>
  )
}

export default ResultEpisodeItem
