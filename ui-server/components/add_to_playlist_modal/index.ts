import { getCurrentUserPlaylists } from 'actions/playlist'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { makeGetUserPlaylists } from 'selectors/entities/playlists'
import { getCurrentUserId } from 'selectors/entities/users'
import { AppState } from 'store'
import { AppActions, CLOSE_MODAL, SHOW_CREATE_PLAYLIST_MODAL } from 'types/actions'
import AddToPlaylistModal, { DispatchToProps, OwnProps, StateToProps } from './add_to_playlist_modal'

function makeMapStateToProps() {
  const getUserPlaylists = makeGetUserPlaylists()

  return (state: AppState): StateToProps => ({
    playlists: getUserPlaylists(state, getCurrentUserId(state)),
    isLoading: false
  })
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    showCreatePlaylistModal: () =>
      dispatch({ type: SHOW_CREATE_PLAYLIST_MODAL }),
    loadPlaylists: bindActionCreators(getCurrentUserPlaylists, dispatch),
    closeModal: () => dispatch({ type: CLOSE_MODAL }),
  }
}

export default connect<StateToProps, DispatchToProps, OwnProps, AppState>(
  makeMapStateToProps(),
  mapDispatchToProps,
)(AddToPlaylistModal)
