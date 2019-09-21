import { connect } from 'react-redux'
import { AppState } from 'store'
import TrendingPage, { OwnProps, StateToProps } from './trending'

function mapStateToProps(state: AppState): StateToProps {
  return {
    reqState: state.requests.podcast.getTrendingPodcasts,
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  TrendingPage,
)
