import { ChartLink } from 'components/link'
import React from 'react'
import { Chart } from 'types/app'

export interface StateToProps {
  relatedCharts: Chart[]
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
    <div className="py-4 px-3 border border-gray-500 rounded-xl">
      <h2 className="text-lg text-gray-700 px-2">{'More related charts'}</h2>
      <hr className="mt-2 mb-4 border-gray-400" />
      {relatedCharts.map((c) => (
        <ChartLink chartId={c.id}>
          <a>
            <li
              key={c.id}
              className="block flex items-center h-10 pl-6 rounded-full hover:bg-gray-200"
            >
              <span className="text-gray-900">{c.title}</span>
              &nbsp;&nbsp;
              <span className="text-sm text-gray-700 tracking-wider">
                {c.subTitle}
              </span>
            </li>
          </a>
        </ChartLink>
      ))}
    </div>
  )
}

export default RelatedCharts
