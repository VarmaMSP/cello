import ModalContainer from 'components/modal/modal_container'
import Overlay from 'components/modal/overlay'
import React from 'react'
import { Playlist } from 'types/models'
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
    <Overlay background="rgba(0, 0, 0, 0.65)">
      <ModalContainer className="modal-slim" header="Add to Playlist">
        <div className="h-full flex flex-col">
          {playlists.length === 0 ? (
            <div className="flex-1 mt-5 text-center tracking-wider text-gray-800">
              {'No playlists found, create one.'}
            </div>
          ) : (
            <div className="flex-1 overflow-y-auto">
              {playlists.map((playlist) => (
                <PlaylistsListItem
                  key={playlist.id}
                  playlist={playlist}
                  episodeId={episodeId}
                />
              ))}
            </div>
          )}

          <button
            className="block flex-none w-full text-sm font-medium text-center text-purple-700 py-1 mr-6 border-2 border-purple-600 rounded-lg"
            onClick={() => showCreatePlaylistModal()}
          >
            NEW PLAYLIST
          </button>
        </div>
      </ModalContainer>
    </Overlay>
  )
}

export default AddToPlaylistModal
