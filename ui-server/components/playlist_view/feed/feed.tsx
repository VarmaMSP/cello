import React from 'react'
import { Playlist } from 'types/app'

export interface StateToProps {
  playlists: Playlist[]
  receviedAll: boolean
  isLoadingMore: boolean
}

export interface DispatchToProps {
  loadMore: (offset: number) => void
}

const Feed: React.FC<StateToProps & DispatchToProps> = ({ playlists }) => {
  return (
    <>
      {playlists.map((p) => (
        <div>{p.title}</div>
      ))}
    </>
  )
}

export default Feed
