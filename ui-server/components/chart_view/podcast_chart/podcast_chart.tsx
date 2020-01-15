import { PodcastLink } from 'components/link'
import React from 'react'
import { Curation, Podcast } from 'types/app'
import { getImageUrl } from 'utils/dom'

export interface StateToProps {
  chart: Curation
  podcasts: Podcast[]
}

export interface OwnProps {
  chartId: string
}

const PodcastChart: React.FC<StateToProps & OwnProps> = ({
  chart,
  podcasts,
}) => {
  return (
    <div>
      <h2 className="text-xl text-gray-700">{chart.title}</h2>
      <hr className="mt-1 mb-4 border-gray-400" />
      <ol>
        {podcasts.map((p, i) => (
          <PodcastLink key={p.id} podcastUrlParam={p.urlParam}>
            <a>
              <li className="block flex items-center py-2 md:px-2 my-3 md:hover:bg-gray-200 rounded-lg">
                <div className="w-6 flex-none md:ml-1 md:mr-4 mr-2 text-sm text-gray-600">
                  {`${i + 1}.`}
                </div>
                <img
                  className="w-20 h-20 mr-4 flex-none object-contain rounded border"
                  src={getImageUrl(p.urlParam)}
                />
                <div>
                  <div className="mb-2 font-medium text-gray-900 tracking-wide leading-tight line-clamp-1">
                    {p.title}
                  </div>
                  <div className="text-xs text-gray-700 tracking-wider line-clamp-2">
                    {p.summary}
                  </div>
                </div>
              </li>
            </a>
          </PodcastLink>
        ))}
      </ol>
    </div>
  )
}

export default PodcastChart
