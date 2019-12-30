import React from 'react'
import { Playlist } from 'types/app'

export interface StateToProps {
  playlist: Playlist
}

export interface OwnProps {
  playlistId: string
}

const PlaylistView: React.FC<StateToProps & OwnProps> = () => {
  return <div></div>
}

export default PlaylistView
