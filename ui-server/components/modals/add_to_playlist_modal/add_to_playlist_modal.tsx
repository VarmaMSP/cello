import React, { useEffect, useState } from 'react'
import { RequestState } from 'reducers/requests/utils'
import { Playlist } from 'types/app'
import ModalContainer from '../components/modal_container'
import Overlay from '../components/overlay'
import Select from './components/select'

export interface StateToProps {
  playlists: Playlist[]
  reqState: RequestState
}

export interface DispatchToProps {
  loadPlaylists: () => void
  closeModal: () => void
}

export interface OwnProps {
  skipLoad?: boolean
}

const AddToPlaylistModal: React.FC<
  StateToProps & DispatchToProps & OwnProps
> = ({ playlists, closeModal, loadPlaylists, reqState }) => {
  const [selected, setSelected] = useState<string>('')

  useEffect(() => {
    loadPlaylists()
  }, [])

  return (
    <Overlay background="rgba(0, 0, 0, 0.8)">
      <ModalContainer handleClose={closeModal} closeUponClicking="CROSS">
        <div className="flex flex-col h-full justify-around">
          <h4 className="flex-none block text-lg mb-4">{'Add to Playlist'}</h4>
          <div className="flex-auto">
            {reqState.status === 'STARTED' ? (
              'LOADING ...'
            ) : (
              <Select
                options={playlists.map((p) => ({ id: p.id, display: p.title }))}
                selected={selected}
                handleSelect={setSelected}
              />
            )}
          </div>
          <div className="flex flex-none justify-end items-center mb-4">
            <div className="w-32 text-sm font-medium text-center text-purple-400 py-1 mr-6 border-2 border-purple-400 rounded-lg">
              New Playlist
            </div>
            <div className="w-32 px-4 py-1 text-sm text-center text-gray-100  bg-purple-500 rounded-lg">
              Add
            </div>
          </div>
        </div>
      </ModalContainer>
    </Overlay>
  )
}

export default AddToPlaylistModal
