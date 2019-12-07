import GridResponsive from 'components/grid_responsive'
import React from 'react'
import { Podcast } from 'types/app'
import { getImageUrl } from 'utils/dom'

export interface StateToProps {
  recommended: Podcast[]
}

const Recommended: React.FC<StateToProps> = ({ recommended }) => {
  return (
    <div className="mb-6">
      <h2 className="text-xl text-gray-700">{'Trending'}</h2>
      <hr className="mt-1 mb-4" />
      <GridResponsive cols={{ SM: 4, MD: 5, LG: 8 }}>
        {recommended.map((podcast) => (
          <div key={podcast.id} className="px-1 mb-4">
            <img
              className="w-full h-auto mb-2 flex-none object-contain rounded-lg border"
              src={getImageUrl(podcast.urlParam)}
            />
            <p className="md:text-xs text-2xs text-gray-800 tracking-wide leading-snug md:mb-1 line-clamp-1">
              {podcast.title}
            </p>
            <p className="md:text-xs text-2xs text-gray-600 leading-snug line-clamp-1">
              {podcast.author}
            </p>
          </div>
        ))}
      </GridResponsive>
    </div>
  )
}

export default Recommended
