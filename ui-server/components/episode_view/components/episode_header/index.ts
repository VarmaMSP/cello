import { loadAndShowAddToPlaylistModal } from 'actions/playlist'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getPodcastById } from 'selectors/entities/podcasts'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import EpisodeHeader, { DispatchToProps, OwnProps, StateToProps } from './episode_header'

function mapStateToProps(state: AppState, { episode }: OwnProps): StateToProps {
  return {
    podcast: getPodcastById(state, episode.podcastId),
  }
}

function mapDispatchToProps(
  dispatch: Dispatch<AppActions>,
  { episode }: OwnProps,
): DispatchToProps {
  return {
    showAddToPlaylistModal: () =>
      bindActionCreators(loadAndShowAddToPlaylistModal, dispatch)(episode.id),
  }
}

export default connect<StateToProps, DispatchToProps, OwnProps, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(EpisodeHeader)
