import PageLayout from 'components/page_layout'
import React from 'react'
import PodcastChart from './podcast_chart'
import RelatedCharts from './related_charts'

interface OwnProps {
  chartId: string
}

const ChartView: React.FC<OwnProps> = ({ chartId }) => {
  return (
    <PageLayout>
      <PodcastChart chartId={chartId} />
      <RelatedCharts chartId={chartId} />
    </PageLayout>
  )
}

export default ChartView
