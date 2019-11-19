import { createPlaylist } from 'actions/playlist'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { AppState } from 'store'
import { AppActions, CLOSE_MODAL } from 'types/actions'
import CreatePlaylistModal, {
  DispatchToProps,
  StateToProps,
} from './create_playlist_modal'

function mapStateToProps(): StateToProps {
  return {
    isLoading: false,
  }
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    closeModal: () => dispatch({ type: CLOSE_MODAL }),
    createPlaylist: (
      title: string,
      privacy: 'PUBLIC' | 'PRIVATE' | 'ANONYMOUS',
    ) => bindActionCreators(createPlaylist, dispatch)(title, privacy),
  }
}

export default connect<StateToProps, DispatchToProps, {}, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(CreatePlaylistModal)
