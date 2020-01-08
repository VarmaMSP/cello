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
      <div className="flex-none mr-1">
        <img
          className="md:w-28 w-16 md:h-28 w-16 object-contain rounded-lg border cursor-default"
          src={getImageUrl(podcast.urlParam)}
        />
      </div>

      <div className="md:pl-4 pl-1">
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

        <PodcastLink podcastUrlParam={podcast.urlParam}>
          <a
            className="text-sm text-grey-800 hover:text-black tracking-wide line-clamp-1"
            style={{ margin: '3px 0px' }}
            dangerouslySetInnerHTML={{
              __html:
                (podcastSearchResult && podcastSearchResult.author) ||
                podcast.author,
            }}
          />
        </PodcastLink>

        <PodcastLink podcastUrlParam={podcast.urlParam}>
          <a
            className="mt-1 text-xs text-gray-700 leading-snug tracking-wider line-clamp-2"
            style={{ hyphens: 'auto' }}
            dangerouslySetInnerHTML={{
              __html: podcastSearchResult.description || podcast.summary,
            }}
          />
        </PodcastLink>
      </div>
    </div>
  )
}

export default PodcastPreview
