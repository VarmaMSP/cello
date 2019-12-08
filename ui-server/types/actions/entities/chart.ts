import { Chart, Podcast } from 'types/app'

export const RECEIVED_PODCAST_CHART = 'RECEIVED_PODCAST_CHART'
export const RECEIVED_PODCAST_CHARTS = 'RECEIVED_PODCAST_CHARTS'
export const RECEIVED_CHART_PODCASTS = 'RECEIVED_CHART_PODCASTS'

export interface ReceivedPodcastChartAction {
  type: typeof RECEIVED_PODCAST_CHART
  chart: Chart
}

export interface ReceivedPodcastChartsAction {
  type: typeof RECEIVED_PODCAST_CHARTS
  charts: Chart[]
}

export interface ReceivedChartPodcastsAction {
  type: typeof RECEIVED_CHART_PODCASTS
  chartId: string
  podcasts: Podcast[]
}

export type ChartActionTypes =
  | ReceivedPodcastChartAction
  | ReceivedPodcastChartsAction
  | ReceivedChartPodcastsAction
