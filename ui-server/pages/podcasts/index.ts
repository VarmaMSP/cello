import { Dispatch, bindActionCreators } from 'redux'
import { getPodcast } from '../../actions/podcast'
import PodcastPage, { OwnProps, StateToProps, DispatchToProps } from './podcast'
import { connect } from 'react-redux'
import { AppState } from 'store'

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
