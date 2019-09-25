import { connect } from 'react-redux'
import { makeGetTrendingPodcasts } from 'selectors/entities/podcasts'
import { AppState } from 'store'
import ListTrendingPodcasts, { StateToProps } from './list_trending_podcasts'

function makeMapStateToProps() {
  const getTrendingPodcasts = makeGetTrendingPodcasts()
  return (state: AppState) => ({
    trendingPodcasts: getTrendingPodcasts(state),
  })
}

export default connect<StateToProps, {}, {}, AppState>(makeMapStateToProps())(
  ListTrendingPodcasts,
)
