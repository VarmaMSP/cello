import ModalContainer from 'components/modal/modal_container'
import Overlay from 'components/modal/overlay'
import { Formik } from 'formik'
import React from 'react'
import { PlaylistPrivacy } from 'types/app'

export interface StateToProps {
  isLoading: boolean
}

export interface DispatchToProps {
  createPlaylist: (title: string, privacy: PlaylistPrivacy) => void
}

export interface OwnProps {
  episodeId: string
}

type Props = StateToProps & DispatchToProps & OwnProps

const CreatePlaylistModal: React.FC<Props> = ({ createPlaylist }) => {
  return (
    <Overlay background="rgba(0, 0, 0, 0.61)">
      <ModalContainer header="Create Playlist">
        <div className="flex flex-col h-full">
          <Formik
            initialValues={{ title: '', privacy: 'PUBLIC' }}
            validate={() => ({})}
            onSubmit={(values) => {
              createPlaylist(values.title, values.privacy as PlaylistPrivacy)
            }}
          >
            {({ values, isSubmitting, handleChange, handleSubmit }) => (
              <form
                className="block flex-1 flex flex-col h-full"
                onSubmit={handleSubmit}
              >
                <div className="flex-1">
                  <label className="block">
                    <span className="text-gray-700">Title</span>
                    <input
                      type="text"
                      name="title"
                      onChange={handleChange}
                      className="form-input block md:w-2/3 w-full mt-2"
                      placeholder="title"
                      value={values.title}
                    />
                  </label>

                  <label className="block mt-4">
                    <span className="text-gray-700">Privacy</span>
                    <select className="form-select block md:w-2/3 w-full mt-2">
                      <option value="PUBLIC">Public</option>
                      <option value="PRIVATE">Private</option>
                    </select>
                  </label>
                </div>

                <div className="relative flex-none h-6 my-4">
                  <button
                    type="submit"
                    className="block absolute right-0 top-0 md:w-32 w-full py-1 text-sm text-center text-gray-100 bg-purple-500 rounded-lg"
                    disabled={isSubmitting}
                  >
                    Submit
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

export default CreatePlaylistModal
