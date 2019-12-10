import { addEpisodeToPlaylists } from 'actions/playlist'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { makeGetUserPlaylists } from 'selectors/entities/playlists'
import { getCurrentUserId } from 'selectors/entities/users'
import { AppState } from 'store'
import { AppActions, SHOW_CREATE_PLAYLIST_MODAL } from 'types/actions'
import AddToPlaylistModal, { DispatchToProps, OwnProps, StateToProps } from './add_to_playlist_modal'

function makeMapStateToProps() {
  const getUserPlaylists_ = makeGetUserPlaylists()

  return (state: AppState): StateToProps => ({
    playlists: getUserPlaylists_(state, getCurrentUserId(state)),
    isLoading: false,
  })
}

function mapDispatchToProps(
  dispatch: Dispatch<AppActions>,
  { episodeId }: OwnProps,
): DispatchToProps {
  return {
    showCreatePlaylistModal: () =>
      dispatch({ type: SHOW_CREATE_PLAYLIST_MODAL, episodeId }),
    addEpisodeToPlaylists: (playlistIds: string[]) =>
      bindActionCreators(addEpisodeToPlaylists, dispatch)(
        episodeId,
        playlistIds,
      ),
  }
}

export default connect<StateToProps, DispatchToProps, OwnProps, AppState>(
  makeMapStateToProps(),
  mapDispatchToProps,
)(AddToPlaylistModal)
