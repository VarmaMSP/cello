import { connect } from 'react-redux'
import { Dispatch } from 'redux'
import { AppState } from 'store'
import { AppActions, CLOSE_MODAL } from 'types/actions'
import CreatePlaylistModal, {
  DispatchToProps,
  StateToProps,
} from './create_playlist_modal'

function mapStateToProps(state: AppState): StateToProps {
  return {
    reqState: state.requests.playlist.createPlaylist,
  }
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    closeModal: () => dispatch({ type: CLOSE_MODAL }),
    showAddToPlaylistModal: () => dispatch({ type: CLOSE_MODAL }),
  }
}

export default connect<StateToProps, DispatchToProps, {}, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(CreatePlaylistModal)
