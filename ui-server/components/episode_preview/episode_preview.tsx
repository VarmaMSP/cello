import classnames from 'classnames'
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
  small?: boolean
  showIcon?: boolean
}

const EpisodePreview: React.FC<StateToProps & OwnProps> = ({
  episode,
  podcast,
  episodeSearchResult,
  small = false,
  showIcon = false,
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
              'md:text-base text-sm font-medium tracking-wide font-normal line-clamp-2',
              { 'mb-1': small },
            )}
            dangerouslySetInnerHTML={{
              __html:
                (episodeSearchResult && episodeSearchResult.title) ||
                episode.title,
            }}
          />
        </EpisodeLink>

        {!small && (
          <PodcastLink podcastUrlParam={podcast.urlParam}>
            <a className="md:text-sm text-xs text-grey-800 mb-2 tracking-wide line-clamp-1">
              {podcast.title}
            </a>
          </PodcastLink>
        )}

        <div className="text-xs md:break-normal break-all tracking-wide leading-sung md:line-clamp-2 line-clamp-3 cursor-default">
          <span className="text-teal-800">{`${formatDistanceToNow(
            parseISO(episode.pubDate),
          )} ago`}</span>
          <span className="mx-2 text-black font-extrabold">&middot;</span>
          <span
            className="text-gray-800"
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
