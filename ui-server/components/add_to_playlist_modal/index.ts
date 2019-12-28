import { connect } from 'react-redux'
import { Dispatch } from 'redux'
import { makeGetPlaylistsByUser } from 'selectors/entities/playlists'
import { getSignedInUserId } from 'selectors/session'
import { AppState } from 'store'
import {
  AppActions,
  MODAL_MANAGER_SHOW_CREATE_PLAYLIST_MODAL,
} from 'types/actions'
import AddToPlaylistModal, {
  DispatchToProps,
  OwnProps,
  StateToProps,
} from './add_to_playlist_modal'

function makeMapStateToProps() {
  const getPlaylistsByUser = makeGetPlaylistsByUser()

  return (state: AppState): StateToProps => ({
    playlists: getPlaylistsByUser(state, getSignedInUserId(state)),
  })
}

function mapDispatchToProps(
  dispatch: Dispatch<AppActions>,
  { episodeId }: OwnProps,
): DispatchToProps {
  return {
    showCreatePlaylistModal: () =>
      dispatch({
        type: MODAL_MANAGER_SHOW_CREATE_PLAYLIST_MODAL,
        episodeId,
      }),
    addEpisodeToPlaylists: (_: string[]) => {},
  }
}

export default connect<StateToProps, DispatchToProps, OwnProps, AppState>(
  makeMapStateToProps(),
  mapDispatchToProps,
)(AddToPlaylistModal)
