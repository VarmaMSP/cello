import GridResponsive from 'components/grid_responsive'
import { PodcastLink } from 'components/link'
import React from 'react'
import { Podcast } from 'types/app'
import { getImageUrl } from 'utils/dom'

export interface StateToProps {
  podcasts: Podcast[]
}

const Recommended: React.FC<StateToProps> = ({ podcasts }) => {
  return (
    <div className="mb-6">
      <h2 className="text-xl text-gray-700">{'Trending'}</h2>
      <hr className="mt-1 mb-4 border-gray-400" />
      <GridResponsive cols={{ SM: 4, MD: 5, LG: 8 }}>
        {podcasts.map((p) => (
          <div key={p.id} className="flex-none px-1 mb-4">
            <PodcastLink podcastId={p.id}>
              <a>
                <img
                  className="w-full h-auto mb-2 flex-none object-contain rounded-lg border"
                  src={getImageUrl(p.urlParam)}
                />
              </a>
            </PodcastLink>
            <PodcastLink podcastId={p.id}>
              <a className="md:text-xs text-2xs text-gray-800 tracking-wide leading-snug md:mb-1 line-clamp-1">
                {p.title}
              </a>
            </PodcastLink>
            <PodcastLink podcastId={p.id}>
              <a className="md:text-xs text-2xs text-gray-600 leading-snug line-clamp-1">
                {p.author}
              </a>
            </PodcastLink>
          </div>
        ))}
      </GridResponsive>
    </div>
  )
}

export default Recommended
