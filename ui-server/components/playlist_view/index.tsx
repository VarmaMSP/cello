import NavTabs from 'components/nav_tabs'
import React from 'react'
import { Playlist } from 'models'
import PlaylistHeader from './components/playlist_header'
import HomeTab from './home_tab/home_tab'

export interface OwnProps {
  playlist?: Playlist
  activeTab: string | undefined
}

const PlaylistView: React.FC<OwnProps> = ({ playlist, activeTab }) => {
  if (playlist === undefined) {
    return (
      <div className="mt-8 text-2xl text-gray-900 tracking-wide">
        {'Playlist not found.'}
      </div>
    )
  }

  return (
    <div className="mt-6">
      <PlaylistHeader playlist={playlist} />
      <div className="mt-6 mb-4">
        <NavTabs
          tabs={[
            {
              name: 'playlist',
              pathname: '/playlists',
              query: { playlistUrlParam: playlist.urlParam, skipLoad: true },
              as: `/playlists/${playlist.urlParam}`,
            },
          ]}
          active={activeTab}
          defaultTab="playlist"
        />
      </div>
      <div className="mt-6 mb-4">
        <HomeTab playlist={playlist} />
      </div>
    </div>
  )
}

export default PlaylistView
