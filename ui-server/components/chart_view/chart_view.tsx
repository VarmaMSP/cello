import React from 'react'
import PodcastChart from './podcast_chart'

interface OwnProps {
  chartId: string
}

const ChartView: React.FC<OwnProps> = ({ chartId }) => {
  return (
    <div className="flex md:flex-row flex-col">
      <div className="md:w-2/3 w-fullpr-2">
        <PodcastChart chartId={chartId} />
      </div>
    </div>
  )
}

export default ChartView
