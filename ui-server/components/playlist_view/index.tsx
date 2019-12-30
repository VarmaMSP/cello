import React from 'react'
import { connect } from 'react-redux'
import { getPlaylistById } from 'selectors/entities/playlists'
import { AppState } from 'store'
import { Playlist } from 'types/app'
import PlaylistInfo from './playlist_info'

export interface StateToProps {
  playlist: Playlist
}

export interface OwnProps {
  playlistId: string
}

const PlaylistView: React.FC<StateToProps & OwnProps> = ({ playlist }) => {
  return (
    <div>
      <PlaylistInfo playlist={playlist} />
    </div>
  )
}

function mapStateToProps(
  state: AppState,
  { playlistId }: OwnProps,
): StateToProps {
  return { playlist: getPlaylistById(state, playlistId) }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  PlaylistView,
)
