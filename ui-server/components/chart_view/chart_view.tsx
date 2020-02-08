import { PodcastLink } from 'components/link'
import React from 'react'
import { Curation, Podcast } from 'types/models'
import { getImageUrl } from 'utils/dom'

export interface StateToProps {
  chart: Curation
  podcasts: Podcast[]
}

export interface OwnProps {
  chartId: string
}

const ChartView: React.FC<StateToProps & OwnProps> = ({
  chart,
  podcasts,
}) => {
  return (
    <div>
      <h1 className="pt-3 pb-1 text-xl tracking-wide text-gray-800 font-medium">
        {chart.title}
      </h1>
      <hr className="mb-3" />
      <ol>
        {podcasts.map((p, i) => (
          <PodcastLink key={p.id} podcastUrlParam={p.urlParam}>
            <a className="block py-4 md:px-2 my-4 md:hover:bg-gray-100 rounded-lg">
              <li className="flex">
                <div className="md:block hidden flex-none w-6 ml-1 mr-4 pt-1 text-gray-700">
                  {`${i + 1}.`}
                </div>
                <img
                  className="md:w-22 md:h-22 w-18 h-18 md:mr-4 mr-3 flex-none object-contain rounded-lg border"
                  src={getImageUrl(p.urlParam)}
                />
                <div>
                  <div className="md:mb-0 mb-1 md:text-base text-sm font-medium tracking-wide line-clamp-1">
                    {p.title}
                  </div>
                  <div className="hidden md:text-sm text-xs text-grey-800 mb-2 tracking-wide md:leading-normal leading-relaxed md:line-clamp-1">
                    {p.author}
                  </div>
                  <div
                    className="md:break-normal break-all md:text-xs text-2xs md:text-gray-700 text-gray-800 tracking-wide leading-snug md:line-clamp-2 line-clamp-3"
                    style={{ hyphens: 'auto' }}
                  >
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

export default ChartView
