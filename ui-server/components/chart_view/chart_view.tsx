import PageLayout from 'components/page_layout'
import React from 'react'
import PodcastChart from './podcast_chart'

interface OwnProps {
  chartId: string
}

const ChartView: React.FC<OwnProps> = ({ chartId }) => {
  return (
    <PageLayout>
      <PodcastChart chartId={chartId} />
      <div />
    </PageLayout>
  )
}

export default ChartView
