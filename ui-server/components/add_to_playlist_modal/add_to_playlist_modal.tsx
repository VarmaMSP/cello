import ModalContainer from 'components/modal/modal_container'
import Overlay from 'components/modal/overlay'
import React from 'react'
import { Playlist } from 'types/app'
import PlaylistsListItem from './playlists_list_item'

export interface StateToProps {
  playlists: Playlist[]
}

export interface DispatchToProps {
  showCreatePlaylistModal: () => void
}

export interface OwnProps {
  skipLoad?: boolean
  episodeId: string
}

type Props = StateToProps & DispatchToProps & OwnProps

const AddToPlaylistModal: React.FC<Props> = ({
  playlists,
  episodeId,
  showCreatePlaylistModal,
}) => {
  return (
    <Overlay background="rgba(0, 0, 0, 0.8)">
      <ModalContainer closeUponClicking="CROSS">
        <div className="flex flex-col h-full">
          <h4 className="flex-none block text-lg mb-4">{'Add to Playlist'}</h4>
          <div className="flex-1 overflow-y-auto">
            {playlists.map((playlist) => (
              <PlaylistsListItem
                key={playlist.id}
                playlist={playlist}
                episodeId={episodeId}
              />
            ))}
          </div>
          <div className="flex-none mb-4">
            <button
              className="w-full text-sm font-medium text-center text-purple-400 py-1 mr-6 border-2 border-purple-400 rounded-lg"
              onClick={() => showCreatePlaylistModal()}
            >
              NEW PLAYLIST
            </button>
          </div>
        </div>
      </ModalContainer>
    </Overlay>
  )
}

export default AddToPlaylistModal
