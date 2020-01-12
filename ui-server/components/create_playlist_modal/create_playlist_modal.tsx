import ModalContainer from 'components/modal/modal_container'
import Overlay from 'components/modal/overlay'
import { Formik } from 'formik'
import React from 'react'
import { PlaylistPrivacy } from 'types/app'

export interface StateToProps {
  isLoading: boolean
}

export interface DispatchToProps {
  createPlaylist: (
    title: string,
    privacy: PlaylistPrivacy,
    description: string,
  ) => void
}

export interface OwnProps {
  episodeId: string
}

type Props = StateToProps & DispatchToProps & OwnProps

const CreatePlaylistModal: React.FC<Props> = ({ createPlaylist }) => {
  return (
    <Overlay background="rgba(0, 0, 0, 0.65)">
      <ModalContainer className="modal-slim" header="Create Playlist">
        <Formik
          initialValues={{ title: '', privacy: 'PUBLIC', description: '' }}
          validate={(values) => {
            const errors: Partial<{ title: string; description: string }> = {}

            if (values.title.length <= 2) {
              errors.title = 'Too Short!'
            }
            if (values.title.length > 150) {
              errors.title = 'Too Long!'
            }

            if (values.description.length > 150) {
              errors.description = 'Too long'
            }

            return errors
          }}
          onSubmit={(values) => {
            console.log(values)
            createPlaylist(
              values.title,
              values.privacy as PlaylistPrivacy,
              values.description,
            )
          }}
        >
          {({ values, errors, isSubmitting, handleChange, handleSubmit }) => (
            <form className="flex flex-col h-full" onSubmit={handleSubmit}>
              <div className="flex-1">
                <label className="block">
                  <span className="text-gray-700">Title</span>
                  <span className="ml-4 text-red-600 text-sm font-bold tracking-wide">
                    {errors.title}
                  </span>
                  <input
                    type="text"
                    name="title"
                    onChange={handleChange}
                    className="form-input w-full mt-2"
                    placeholder="title"
                    value={values.title}
                  />
                </label>

                <label className="block mt-4">
                  <span className="text-gray-700">Privacy</span>
                  <select
                    name="privacy"
                    onChange={handleChange}
                    className="form-select w-full mt-2"
                    value={values.privacy}
                  >
                    <option value="PUBLIC">Public</option>
                    <option value="PRIVATE">Private</option>
                  </select>
                </label>

                <label className="block mt-4">
                  <span className="text-gray-700 tracking-wide">
                    {'Description (optional)'}
                  </span>
                  <span className="ml-4 text-red-600 text-sm font-bold tracking-wide">
                    {errors.description}
                  </span>
                  <input
                    type="text"
                    name="description"
                    onChange={handleChange}
                    className="form-input w-full mt-2"
                    placeholder="description"
                    value={values.description}
                  />
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
      </ModalContainer>
    </Overlay>
  )
}

export default CreatePlaylistModal
