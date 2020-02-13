import classnames from 'classnames'
import EpisodeThumbnail from 'components/episode_thumbnail'
import { EpisodeLink, PodcastLink } from 'components/link'
import formatDistanceToNow from 'date-fns/formatDistanceToNow'
import parseISO from 'date-fns/parseISO'
import { Episode, EpisodeSearchResult, Podcast } from 'models'
import React from 'react'

export interface StateToProps {
  episode: Episode
  podcast: Podcast
  episodeSearchResult: EpisodeSearchResult
}

export interface OwnProps {
  episodeId: string
  small?: boolean
  showIcon?: boolean
  showHighlights?: boolean
}

const EpisodePreview: React.FC<StateToProps & OwnProps> = ({
  episode,
  podcast,
  episodeSearchResult,
  small = false,
  showIcon = false,
  showHighlights = false,
}) => {
  return (
    <div className="episode-preview flex md:px-1 py-4 md:hover:bg-gray-100 rounded-lg">
      <div className="flex-none md:mr-4 mr-3">
        <EpisodeThumbnail
          episodeId={episode.id}
          small={small}
          showIcon={showIcon}
        />
      </div>

      <div>
        <EpisodeLink episodeUrlParam={episode.urlParam}>
          <a
            className={classnames(
              'md:text-base text-sm text-black font-semibold tracking-wide line-clamp-2',
              { 'mb-1': small },
            )}
            dangerouslySetInnerHTML={{
              __html: showHighlights
                ? episodeSearchResult.title
                : episode.title,
            }}
          />
        </EpisodeLink>

        {!small && (
          <PodcastLink podcastUrlParam={podcast.urlParam}>
            <a className="md:text-sm text-xs text-gray-900 font-medium mb-2 tracking-wide md:leading-normal leading-relaxed line-clamp-1">
              {podcast.title}
            </a>
          </PodcastLink>
        )}

        <div
          className="text-xs md:break-normal break-all tracking-normal leading-normal md:line-clamp-2 line-clamp-3 cursor-default"
          style={{ hyphens: 'auto' }}
        >
          <span className="text-gray-700 font-medium">{`${formatDistanceToNow(
            parseISO(episode.pubDate),
          )} ago`}</span>
          <span className="mx-2 text-black font-extrabold">&middot;</span>
          <span
            className="text-xs text-gray-800"
            style={{ fontWeight: 450 }}
            dangerouslySetInnerHTML={{
              __html: showHighlights
                ? episodeSearchResult.description
                : episode.summary,
            }}
          />
        </div>
      </div>
    </div>
  )
}

export default EpisodePreview
