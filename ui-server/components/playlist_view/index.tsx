import NavTabs from 'components/nav_tabs'
import React from 'react'
import { connect } from 'react-redux'
import { getPlaylistById } from 'selectors/entities/playlists'
import { AppState } from 'store'
import { Playlist } from 'types/app'
import HomeTab from './home_tab/home_tab'
import PlaylistInfo from './playlist_info'

export interface StateToProps {
  playlist: Playlist
}

export interface OwnProps {
  playlistId: string
  activeTab: string | undefined
}

const PlaylistView: React.FC<StateToProps & OwnProps> = ({
  playlist,
  activeTab,
}) => {
  const playlistUrlParam = playlist.urlParam

  return (
    <div>
      <PlaylistInfo playlist={playlist} />
      <div className="mt-6 mb-4">
        <NavTabs
          tabs={[
            {
              name: 'playlist',
              pathname: '/playlists',
              query: { playlistUrlParam, skipLoad: true },
              as: `/playlists/${playlistUrlParam}`,
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

function mapStateToProps(
  state: AppState,
  { playlistId }: OwnProps,
): StateToProps {
  return { playlist: getPlaylistById(state, playlistId) }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  PlaylistView,
)