export const MODAL_MANAGER_SHOW_SIGN_IN_MODAL =
  'modal_manager/show_sign_in_modal'
export const MODAL_MANAGER_SHOW_ADD_TO_PLAYLIST_MODAL =
  'modal_manager/show_add_to_playlist_modal'
export const MODAL_MANAGER_SHOW_CREATE_PLAYLIST_MODAL =
  'modal_manager/show_create_playlist_modal'
export const MODAL_MANAGER_CLOSE_MODAL = 'modal_manager/close_modal'

export interface ShowSigninModalAction {
  type: typeof MODAL_MANAGER_SHOW_SIGN_IN_MODAL
}

export interface ShowAddToPlaylistModalAction {
  type: typeof MODAL_MANAGER_SHOW_ADD_TO_PLAYLIST_MODAL
  episodeId: string
}

export interface ShowCreatePlaylistModalAction {
  type: typeof MODAL_MANAGER_SHOW_CREATE_PLAYLIST_MODAL
  episodeId: string
}

export interface CloseModalAction {
  type: typeof MODAL_MANAGER_CLOSE_MODAL
}
