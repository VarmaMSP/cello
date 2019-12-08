import React from 'react'
import PodcastChart from './podcast_chart'
import RelatedCharts from './related_charts'

interface OwnProps {
  chartId: string
}

const ChartView: React.FC<OwnProps> = ({ chartId }) => {
  return (
    <div className="flex md:flex-row flex-col">
      <div className="md:w-2/3 w-full pr-4">
        <PodcastChart chartId={chartId} />
      </div>
      <div className="md:w-1/3 w-full pl-2">
        <RelatedCharts chartId={chartId} />
      </div>
    </div>
  )
}

export default ChartView
