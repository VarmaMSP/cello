import { connect } from 'react-redux'
import { getChartById, makeGetPodcastsInChart } from 'selectors/entities/charts'
import { AppState } from 'store'
import PodcastList, { OwnProps, StateToProps } from './podcast_chart'

function makeMapStateToProps() {
  const getPodcastsInChart = makeGetPodcastsInChart()

  return (state: AppState, { chartId }: OwnProps): StateToProps => {
    return {
      chart: getChartById(state, chartId),
      podcasts: getPodcastsInChart(state, chartId),
    }
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(
  makeMapStateToProps(),
)(PodcastList)
