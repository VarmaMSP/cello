import { connect } from 'react-redux'
import { makeGetPodcastsInChart } from 'selectors/entities/charts'
import { AppState } from 'store'
import Recommended, { StateToProps } from './recommended'

function makeMapStateToProps() {
  const getPodcastsInChart = makeGetPodcastsInChart()

  return (state: AppState): StateToProps => {
    return { podcasts: getPodcastsInChart(state, 'recommended')}
  }
}

export default connect<StateToProps, {}, {}, AppState>(makeMapStateToProps())(
  Recommended,
)
