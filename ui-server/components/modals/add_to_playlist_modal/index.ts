import { getCurrentUserPlaylists } from 'actions/playlist'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { makeGetUserPlaylists } from 'selectors/entities/playlists'
import { getCurrentUserId } from 'selectors/entities/users'
import { AppState } from 'store'
import { AppActions, CLOSE_MODAL } from 'types/actions'
import AddToPlaylistModal, {
  DispatchToProps,
  OwnProps,
  StateToProps,
} from './add_to_playlist_modal'

function makeMapStateToProps() {
  const getUserPlaylists = makeGetUserPlaylists()

  return (state: AppState): StateToProps => ({
    playlists: getUserPlaylists(state, getCurrentUserId(state)),
    reqState: state.requests.playlist.getUserPlaylists,
  })
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    loadPlaylists: bindActionCreators(getCurrentUserPlaylists, dispatch),
    closeModal: () => dispatch({ type: CLOSE_MODAL }),
  }
}

export default connect<StateToProps, DispatchToProps, OwnProps, AppState>(
  makeMapStateToProps(),
  mapDispatchToProps,
)(AddToPlaylistModal)
