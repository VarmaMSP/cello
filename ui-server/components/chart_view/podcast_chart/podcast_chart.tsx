import { PodcastLink } from 'components/link'
import React from 'react'
import { Chart, Podcast } from 'types/app'
import { getImageUrl } from 'utils/dom'

export interface StateToProps {
  chart: Chart
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
          <PodcastLink key={p.id} podcastId={p.id}>
            <a>
              <li className="block flex items-center py-2 my-1 hover:bg-gray-200 rounded-lg">
                <div className="w-6 md:ml-1 md:mr-4 mr-2 text-sm text-gray-600">{`${i + 1}.`}</div>
                <img
                  className="w-16 h-16 mr-4 flex-none object-contain rounded border cursor-default"
                  src={getImageUrl(p.urlParam)}
                />
                <div className="h-16">
                  <div className="text-gray-900 tracking-wide leading-loose line-clamp-1">
                    {p.title}
                  </div>
                  <div className="text-sm text-gray-800 traking-wide leading-relaxed line-clamp-1">
                    {p.author}
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
