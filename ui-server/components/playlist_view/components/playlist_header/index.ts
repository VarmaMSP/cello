import { deletePlaylist } from 'actions/playlist'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getSignedInUserId } from 'selectors/session'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import PlaylistHeader, { DispatchToProps, OwnProps, StateToProps } from './playlist_header'

function mapStateToProps(state: AppState): StateToProps {
  return {
    signedInUserId: getSignedInUserId(state)
  }
}

function mapDispatchToProps(
  dispatch: Dispatch<AppActions>,
  { playlist }: OwnProps,
): DispatchToProps {
  return {
    removePlaylist: () =>
      bindActionCreators(deletePlaylist, dispatch)(playlist.id),
  }
}

export default connect<StateToProps, DispatchToProps, OwnProps, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(PlaylistHeader)
