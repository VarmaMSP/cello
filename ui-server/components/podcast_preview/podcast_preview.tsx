import { PodcastLink } from 'components/link'
import { Podcast, PodcastSearchResult } from 'types/app'
import { getImageUrl } from 'utils/dom'

export interface StateToProps {
  podcast: Podcast
  podcastSearchResult: PodcastSearchResult
}

export interface OwnProps {
  podcastId: string
}

const PodcastPreview: React.FC<StateToProps & OwnProps> = ({
  podcast,
  podcastSearchResult,
}) => {
  return (
    <div className="flex mb-14">
      <div className="flex-none md:mr-4 mr-3">
        <img
          className="md:w-28 w-22 md:h-28 w-22 object-contain rounded-lg border cursor-default"
          src={getImageUrl(podcast.urlParam)}
        />
      </div>

      <div>
        <PodcastLink podcastUrlParam={podcast.urlParam}>
          <a
            className="md:text-base text-sm font-medium tracking-wide line-clamp-2"
            dangerouslySetInnerHTML={{
              __html:
                (podcastSearchResult && podcastSearchResult.title) ||
                podcast.title,
            }}
          />
        </PodcastLink>

        <div
          className="md:text-sm text-xs text-grey-800 mb-2 tracking-wide md:leading-normal leading-relaxed line-clamp-1"
          dangerouslySetInnerHTML={{
            __html:
              (podcastSearchResult && podcastSearchResult.author) ||
              podcast.author,
          }}
        />

        <div
          className="text-xs text-gray-800 md:break-normal break-all leading-snug tracking-wide md:line-clamp-2 line-clamp-3 cursor-default"
          style={{ hyphens: 'auto' }}
          dangerouslySetInnerHTML={{
            __html: podcastSearchResult.description || podcast.summary,
          }}
        />
      </div>
    </div>
  )
}

export default PodcastPreview
