import { connect } from 'react-redux'
import { getPlaylistById } from 'selectors/entities/playlists'
import { AppState } from 'store'
import PlaylistView, { OwnProps, StateToProps } from './playlist_view'

function mapStateToProps(
  state: AppState,
  { playlistId }: OwnProps,
): StateToProps {
  return { playlist: getPlaylistById(state, playlistId) }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  PlaylistView,
)
