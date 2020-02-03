import { ChartLink } from 'components/link'
import React from 'react'
import { Curation } from 'types/models'

export interface StateToProps {
  relatedCharts: Curation[]
}

export interface OwnProps {
  chartId: string
}

const RelatedCharts: React.FC<StateToProps & OwnProps> = ({
  relatedCharts,
}) => {
  if (relatedCharts.length === 0) {
    return <></>
  }

  return (
    <div className="px-3 pt-3 pb-2 rounded-xl bg-gray-200">
      <h2 className="mb-3 text text-gray-800 tracking-wide">
        {'Related charts'}
      </h2>
      {relatedCharts.map((c) => (
        <ChartLink key={c.id} chartId={c.id}>
          <a>
            <li
              key={c.id}
              className="block flex items-center h-10 pl-6 rounded-full hover:bg-gray-200"
            >
              <span className="text-gray-900">{c.title}</span>
            </li>
          </a>
        </ChartLink>
      ))}
    </div>
  )
}

export default RelatedCharts
