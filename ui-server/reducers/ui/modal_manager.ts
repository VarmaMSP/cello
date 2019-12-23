import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { Modal } from 'types/app'

const activeModal: Reducer<Modal, T.AppActions> = (
  state = { type: 'NONE' },
  action,
) => {
  switch (action.type) {
    case T.MODAL_MANAGER_SHOW_SIGN_IN_MODAL:
      return { type: 'SIGNIN_MODAL' }
    case T.MODAL_MANAGER_SHOW_ADD_TO_PLAYLIST_MODAL:
      return { type: 'ADD_TO_PLAYLIST_MODAL', episodeId: action.episodeId }
    case T.MODAL_MANAGER_SHOW_CREATE_PLAYLIST_MODAL:
      return { type: 'CREATE_PLAYLIST_MODAL', episodeId: action.episodeId }
    case T.MODAL_MANAGER_CLOSE_MODAL:
      return { type: 'NONE' }
    default:
      return state
  }
}

export default combineReducers({
  activeModal,
})
