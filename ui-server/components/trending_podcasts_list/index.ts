import { connect } from 'react-redux'
import { makeGetTrendingPodcasts } from 'selectors/entities/podcasts'
import { AppState } from 'store'
import TrendingPodcastsList, { StateToProps } from './trending_podcasts_list'

function makeMapStateToProps() {
  const getTrendingPodcasts = makeGetTrendingPodcasts()
  return (state: AppState) => ({
    trendingPodcasts: getTrendingPodcasts(state),
  })
}

export default connect<StateToProps, {}, {}, AppState>(makeMapStateToProps())(
  TrendingPodcastsList,
)
