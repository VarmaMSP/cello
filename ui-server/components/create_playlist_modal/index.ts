import { createPlaylist } from 'actions/playlist'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getCreatePlaylistStatus } from 'selectors/request'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import { PlaylistPrivacy } from 'types/app'
import CreatePlaylistModal, {
  DispatchToProps,
  OwnProps,
  StateToProps,
} from './create_playlist_modal'

function mapStateToProps(state: AppState): StateToProps {
  return {
    isLoading: getCreatePlaylistStatus(state) === 'IN_PROGRESS',
  }
}

function mapDispatchToProps(
  dispatch: Dispatch<AppActions>,
  { episodeId }: OwnProps,
): DispatchToProps {
  return {
    createPlaylist: (title: string, privacy: PlaylistPrivacy) =>
      bindActionCreators(createPlaylist, dispatch)(title, privacy, episodeId),
  }
}

export default connect<StateToProps, DispatchToProps, OwnProps, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(CreatePlaylistModal)
