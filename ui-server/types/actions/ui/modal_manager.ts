export const MODAL_MANAGER_SHOW_SIGN_IN_MODAL =
  'modal_manager/show_sign_in_modal'
export const MODAL_MANAGER_SHOW_ADD_TO_PLAYLIST_MODAL =
  'modal_manager/show_add_to_playlist_modal'
export const MODAL_MANAGER_SHOW_CREATE_PLAYLIST_MODAL =
  'modal_manager/show_create_playlist_modal'
export const MODAL_MANAGER_CLOSE_MODAL = 'modal_manager/close_modal'

interface ShowSigninModalAction {
  type: typeof MODAL_MANAGER_SHOW_SIGN_IN_MODAL
}

interface ShowAddToPlaylistModalAction {
  type: typeof MODAL_MANAGER_SHOW_ADD_TO_PLAYLIST_MODAL
  episodeId: string
}

interface ShowCreatePlaylistModalAction {
  type: typeof MODAL_MANAGER_SHOW_CREATE_PLAYLIST_MODAL
  episodeId: string
}

interface CloseModalAction {
  type: typeof MODAL_MANAGER_CLOSE_MODAL
}

export type ModalManagerActionTypes =
  | ShowSigninModalAction
  | ShowAddToPlaylistModalAction
  | ShowCreatePlaylistModalAction
  | CloseModalAction
