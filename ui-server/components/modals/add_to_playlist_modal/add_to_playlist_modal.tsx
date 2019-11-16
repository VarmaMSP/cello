import { Formik } from 'formik'
import React, { useEffect } from 'react'
import { RequestState } from 'reducers/requests/utils'
import { Playlist } from 'types/app'
import ModalContainer from '../components/modal_container'
import Overlay from '../components/overlay'

export interface StateToProps {
  playlists: Playlist[]
  reqState: RequestState
}

export interface DispatchToProps {
  showCreatePlaylistModal: () => void
  loadPlaylists: () => void
  closeModal: () => void
}

export interface OwnProps {
  skipLoad?: boolean
  episodeId: string
}

const AddToPlaylistModal: React.FC<StateToProps &
  DispatchToProps &
  OwnProps> = ({
  playlists,
  closeModal,
  showCreatePlaylistModal,
  loadPlaylists,
  reqState,
}) => {
  useEffect(() => {
    loadPlaylists()
  }, [])

  return (
    <Overlay background="rgba(0, 0, 0, 0.8)">
      <ModalContainer handleClose={closeModal} closeUponClicking="CROSS">
        <div className="flex flex-col h-full">
          <h4 className="flex-none block text-lg mb-4">{'Add to Playlist'}</h4>
          {reqState.status === 'STARTED' ? (
            'LOADING ...'
          ) : (
            <Formik
              initialValues={{ playlistIds: [] as string[] }}
              validate={() => ({})}
              onSubmit={(values) => {
                console.log(values)
              }}
            >
              {({ values, isSubmitting, handleSubmit, setFieldValue }) => (
                <form
                  className="block flex-1 flex flex-col h-full"
                  onSubmit={handleSubmit}
                >
                  <div className="flex-1">
                    {playlists.map(({ id, title }) => (
                      <div className="mb-1">
                        <label className="inline-flex items-center">
                          <input
                            type="checkbox"
                            name="playlistIds"
                            className="form-checkbox"
                            checked={values.playlistIds.includes(id)}
                            onChange={(e) =>
                              e.target.checked
                                ? setFieldValue('playlistIds', [
                                    ...new Set([...values.playlistIds, id]),
                                  ])
                                : setFieldValue(
                                    'playlistIds',
                                    values.playlistIds.filter((x) => x != id),
                                  )
                            }
                          />
                          <span className="ml-2">{title}</span>
                        </label>
                      </div>
                    ))}
                  </div>
                  <div className="flex flex-none md:justify-end justify-center items-center mb-4">
                    <button
                      className="w-32 text-sm font-medium text-center text-purple-400 py-1 mr-6 border-2 border-purple-400 rounded-lg"
                      onClick={() => showCreatePlaylistModal()}
                    >
                      New Playlist
                    </button>
                    <button
                      type="submit"
                      className="w-32 px-4 py-1 text-sm text-center text-gray-100 bg-purple-500 rounded-lg"
                      disabled={isSubmitting}
                    >
                      Add
                    </button>
                  </div>
                </form>
              )}
            </Formik>
          )}
        </div>
      </ModalContainer>
    </Overlay>
  )
}

export default AddToPlaylistModal
