import { connect } from 'react-redux'
import { AppState } from 'store'
import TrendingPage, { StateToProps } from './trending'

function mapStateToProps(state: AppState): StateToProps {
  return {
    reqState: state.requests.podcast.getTrendingPodcasts,
  }
}

export default connect<StateToProps, {}, {}, AppState>(mapStateToProps)(
  TrendingPage,
)
