import { connect } from 'react-redux'
import {
  getCurationById,
  makeGetPodcastsInCuration,
} from 'selectors/entities/curations'
import { AppState } from 'store'
import ChartView, { OwnProps, StateToProps } from './chart_view'

function makeMapStateToProps() {
  const getPodcastsInCuration = makeGetPodcastsInCuration()

  return (state: AppState, { chartId }: OwnProps): StateToProps => {
    return {
      chart: getCurationById(state, chartId),
      podcasts: getPodcastsInCuration(state, chartId),
    }
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(
  makeMapStateToProps(),
)(ChartView)
