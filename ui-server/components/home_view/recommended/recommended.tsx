import Grid from 'components/grid'
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
      <h1 className="pt-3 pb-1 text-xl tracking-wide text-gray-800 font-medium">
        {'Recommended'}
      </h1>
      <hr className="mb-3" />
      <Grid cols={{ SM: 4, MD: 5, LG: 8 }}>
        {podcasts.map((p) => (
          <div key={p.id} className="flex-none px-1 mb-4">
            <PodcastLink podcastUrlParam={p.urlParam}>
              <a>
                <img
                  className="w-full h-auto mb-2 flex-none object-contain rounded-lg border"
                  src={getImageUrl(p.urlParam)}
                />
              </a>
            </PodcastLink>
            <PodcastLink podcastUrlParam={p.urlParam}>
              <a className="md:text-xs text-2xs text-gray-800 tracking-wide leading-snug md:mb-1 line-clamp-1">
                {p.title}
              </a>
            </PodcastLink>
            <PodcastLink podcastUrlParam={p.urlParam}>
              <a className="md:text-xs text-2xs text-gray-600 leading-snug line-clamp-1">
                {p.author}
              </a>
            </PodcastLink>
          </div>
        ))}
      </Grid>
    </div>
  )
}

export default Recommended
