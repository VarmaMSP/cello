import EpisodeThumbnail from 'components/episode_thumbnail'
import { EpisodeLink, PodcastLink } from 'components/link'
import formatDistanceToNow from 'date-fns/formatDistanceToNow'
import parseISO from 'date-fns/parseISO'
import React from 'react'
import { Episode, EpisodeSearchResult, Podcast } from 'types/app'

export interface StateToProps {
  episode: Episode
  podcast: Podcast
  episodeSearchResult: EpisodeSearchResult | undefined
}

export interface OwnProps {
  episodeId: string
}

const EpisodePreview: React.FC<StateToProps & OwnProps> = ({
  episode,
  podcast,
  episodeSearchResult,
}) => {
  return (
    <div className="episode-preview flex mb-6 md:px-2 py-4 md:hover:bg-gray-100 rounded-lg">
      <div className="flex-none mr-2">
        <EpisodeThumbnail episodeId={episode.id} />
      </div>

      <div className="md:pl-4 pl-1">
        <EpisodeLink episodeUrlParam={episode.urlParam}>
          <a
            className="md:text-base text-sm font-medium tracking-wide line-clamp-2"
            dangerouslySetInnerHTML={{
              __html:
                (episodeSearchResult && episodeSearchResult.title) ||
                episode.title,
            }}
          />
        </EpisodeLink>

        <PodcastLink podcastUrlParam={podcast.urlParam}>
          <a className="md:text-sm text-xs text-grey-800 mt-1 mb-2 tracking-wide line-clamp-1">
            {podcast.title}
          </a>
        </PodcastLink>

        <div className="md:text-sm text-2xs md:break-normal break-all tracking-wide line-clamp-3 cursor-default">
          <span className="text-gray-700 font-medium">{`${formatDistanceToNow(
            parseISO(episode.pubDate),
          )} ago`}</span>
          <span className="mx-2 text-black font-extrabold">&middot;</span>
          <span
            className="text-gray-700"
            dangerouslySetInnerHTML={{
              __html:
                (episodeSearchResult && episodeSearchResult.description) ||
                episode.summary,
            }}
          />
        </div>
      </div>
    </div>
  )
}

export default EpisodePreview
