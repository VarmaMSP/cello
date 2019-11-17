import { Formik } from 'formik'
import React from 'react'
import { RequestState } from 'reducers/requests/utils'
import ModalContainer from '../modals/components/modal_container'
import Overlay from '../modals/components/overlay'

export interface StateToProps {
  reqState: RequestState
}

export interface DispatchToProps {
  closeModal: () => void
  createPlaylist: (
    title: string,
    privacy: 'PUBLIC' | 'PRIVATE' | 'ANONYMOUS',
  ) => void
}

const CreatePlaylistModal: React.FC<DispatchToProps> = ({
  closeModal,
  createPlaylist,
}) => {
  return (
    <Overlay background="rgba(0, 0, 0, 0.8)">
      <ModalContainer handleClose={closeModal} closeUponClicking="CROSS">
        <div className="flex flex-col h-full">
          <h4 className="flex-none block text-2xl mb-6">{'Create Playlist'}</h4>
          <Formik
            initialValues={{ title: '', privacy: 'PUBLIC' }}
            validate={() => ({})}
            onSubmit={(values) => {
              createPlaylist(
                values.title,
                values.privacy as 'PUBLIC' | 'PRIVATE',
              )
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
                      <option value="public">Public</option>
                      <option value="private">Private</option>
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
