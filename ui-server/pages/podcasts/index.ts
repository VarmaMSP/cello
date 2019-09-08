import { getPodcast } from 'actions/podcast'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { AppState } from 'store'
import PodcastPage, { DispatchToProps, OwnProps, StateToProps } from './podcast'

function mapStateToProps(state: AppState): StateToProps {
  return {
    reqState: state.requests.podcast.getPodcast,
  }
}

function mapDispatchToProps(dispatch: Dispatch): DispatchToProps {
  return {
    loadPodcast: bindActionCreators(getPodcast, dispatch),
  }
}

export default connect<StateToProps, DispatchToProps, OwnProps, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(PodcastPage)
