import ModalContainer from 'components/modal/modal_container'
import Overlay from 'components/modal/overlay'
import { Formik } from 'formik'
import React from 'react'
import { Playlist } from 'types/app'

export interface StateToProps {
  playlists: Playlist[]
}

export interface DispatchToProps {
  showCreatePlaylistModal: () => void
  addEpisodeToPlaylists: (playlistIds: string[]) => void
}

export interface OwnProps {
  skipLoad?: boolean
  episodeId: string
  closeModal: () => void
}

type Props = StateToProps & DispatchToProps & OwnProps

const AddToPlaylistModal: React.FC<Props> = ({
  playlists,
  closeModal,
  showCreatePlaylistModal,
  addEpisodeToPlaylists,
}) => {
  return (
    <Overlay background="rgba(0, 0, 0, 0.8)">
      <ModalContainer handleClose={closeModal} closeUponClicking="CROSS">
        <div className="flex flex-col h-full">
          <h4 className="flex-none block text-lg mb-4">{'Add to Playlist'}</h4>
          <Formik
            initialValues={{ playlistIds: [] as string[] }}
            validate={() => ({})}
            onSubmit={(values) => {
              addEpisodeToPlaylists(values.playlistIds)
            }}
          >
            {({ values, isSubmitting, handleSubmit, setFieldValue }) => (
              <form
                className="block flex-1 flex flex-col h-full"
                onSubmit={handleSubmit}
              >
                <div className="flex-1 overflow-y-auto">
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
                <div className="flex flex-none md:justify-end justify-center items-center mt-1 mb-4">
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
        </div>
      </ModalContainer>
    </Overlay>
  )
}

export default AddToPlaylistModal
