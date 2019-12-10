import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { Modal } from 'types/app'
import player from './player'

const searchText: Reducer<string, T.AppActions> = (state = '', action) => {
  switch (action.type) {
    case T.SEARCH_BAR_TEXT_CHANGE:
      return action.text
    default:
      return state
  }
}

const showModal: Reducer<Modal, T.AppActions> = (
  state = { type: 'NONE' },
  action,
) => {
  switch (action.type) {
    case T.SHOW_SIGNIN_MODAL:
      return { type: 'SIGNIN_MODAL' }
    case T.SHOW_ADD_TO_PLAYLIST_MODAL:
      return { type: 'ADD_TO_PLAYLIST_MODAL', episodeId: action.episodeId }
    case T.SHOW_CREATE_PLAYLIST_MODAL:
      return { type: 'CREATE_PLAYLIST_MODAL', episodeId: action.episodeId }
    case T.CLOSE_MODAL:
      return { type: 'NONE' }
    default:
      return state
  }
}

export default combineReducers({
  searchText,
  showModal,
  player,
})
