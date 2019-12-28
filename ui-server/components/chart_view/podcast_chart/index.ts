import { connect } from 'react-redux'
import { getCurationById, makeGetPodcastsInCuration } from 'selectors/entities/curations'
import { getPodcastsByIds } from 'selectors/entities/podcasts'
import { AppState } from 'store'
import PodcastList, { OwnProps, StateToProps } from './podcast_chart'

function makeMapStateToProps() {
  const getPodcastsInChart = makeGetPodcastsInCuration()

  return (state: AppState, { chartId }: OwnProps): StateToProps => {
    return {
      chart: getCurationById(state, chartId),
      podcasts: getPodcastsByIds(state, getPodcastsInChart(state, chartId)),
    }
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(
  makeMapStateToProps(),
)(PodcastList)
