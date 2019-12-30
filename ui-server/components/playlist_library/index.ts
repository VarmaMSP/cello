import { connect } from 'react-redux'
import { makeGetPlaylistsByUser } from 'selectors/entities/playlists'
import { getSignedInUserId } from 'selectors/session'
import { AppState } from 'store'
import PlaylistLibrary, { StateToProps } from './playlist_library'

function makeMapStateToProps() {
  const getPlaylistsByUser = makeGetPlaylistsByUser()

  return (state: AppState): StateToProps => {
    return { playlists: getPlaylistsByUser(state, getSignedInUserId(state)) }
  }
}

export default connect<StateToProps, {}, {}, AppState>(makeMapStateToProps())(
  PlaylistLibrary,
)
