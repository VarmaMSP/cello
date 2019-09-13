import { connect } from 'react-redux'
import { AppState } from 'store'
import PodcastPage, { OwnProps, StateToProps } from './podcast'

function mapStateToProps(state: AppState): StateToProps {
  return {
    reqState: state.requests.podcast.getPodcast,
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  PodcastPage,
)
