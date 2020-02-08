import { PodcastLink } from 'components/link'
import { Podcast, PodcastSearchResult } from 'types/models'
import { getImageUrl } from 'utils/dom'

export interface StateToProps {
  podcast: Podcast
  podcastSearchResult: PodcastSearchResult
}

export interface OwnProps {
  podcastId: string
  showHighlights?: boolean
}

const PodcastPreview: React.FC<StateToProps & OwnProps> = ({
  podcast,
  podcastSearchResult,
  showHighlights = false,
}) => {
  return (
    <div className="flex md:px-1 py-4 md:hover:bg-gray-100 rounded-lg">
      <div className="flex-none md:mr-4 mr-3">
        <img
          className="md:w-28 w-22 md:h-28 w-22 object-contain rounded-lg border cursor-default"
          src={getImageUrl(podcast.urlParam)}
        />
      </div>

      <div>
        <PodcastLink podcastUrlParam={podcast.urlParam}>
          <a
            className="md:text-base text-sm font-semibold tracking-wide line-clamp-2"
            dangerouslySetInnerHTML={{
              __html: showHighlights
                ? podcastSearchResult.title
                : podcast.title,
            }}
          />
        </PodcastLink>

        <div
          className="md:text-sm text-xs text-gray-900 font-medium mb-2 tracking-wide md:leading-normal leading-relaxed line-clamp-1"
          dangerouslySetInnerHTML={{
            __html: showHighlights
              ? podcastSearchResult.author
              : podcast.author,
          }}
        />

        <div
          className="text-xs md:break-normal break-all leading-normal tracking-wide md:line-clamp-2 line-clamp-3 cursor-default"
          style={{ hyphens: 'auto' }}
          dangerouslySetInnerHTML={{
            __html: showHighlights
              ? podcastSearchResult.description
              : podcast.summary,
          }}
        />
      </div>
    </div>
  )
}

export default PodcastPreview
