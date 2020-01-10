import { deletePlaylist } from 'actions/playlist'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import PlaylistHeader, { DispatchToProps, OwnProps } from './playlist_header'

function mapDispatchToProps(
  dispatch: Dispatch<AppActions>,
  { playlist }: OwnProps,
): DispatchToProps {
  return {
    removePlaylist: () =>
      bindActionCreators(deletePlaylist, dispatch)(playlist.id),
  }
}

export default connect<{}, DispatchToProps, OwnProps, AppState>(
  null,
  mapDispatchToProps,
)(PlaylistHeader)
