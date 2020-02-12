import { createPlaylist } from 'actions/playlist'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getCreatePlaylistStatus } from 'selectors/request'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import { PlaylistPrivacy } from 'models'
import CreatePlaylistModal, { DispatchToProps, OwnProps, StateToProps } from './create_playlist_modal'

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
    createPlaylist: (title: string, privacy: PlaylistPrivacy, description: string) =>
      bindActionCreators(createPlaylist, dispatch)(
        title,
        privacy,
        description,
        episodeId,
      ),
  }
}

export default connect<StateToProps, DispatchToProps, OwnProps, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(CreatePlaylistModal)
