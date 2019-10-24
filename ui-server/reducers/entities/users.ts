import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { EpisodePlayback, User } from 'types/app'

const currentUserId: Reducer<string, T.AppActions> = (state = '', action) => {
  switch (action.type) {
    case T.RECEIVED_SIGNED_IN_USER:
      return action.user.id
    case T.SIGN_OUT_USER_FORCEFULLY:
    case T.SIGN_OUT_USER_SUCCESS:
      return ''
    default:
      return state
  }
}

const users: Reducer<{ [userId: string]: User }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.RECEIVED_SIGNED_IN_USER:
      return { ...state, [action.user.id]: action.user }
    default:
      return state
  }
}

const playback: Reducer<
  { [episodeId: string]: EpisodePlayback },
  T.AppActions
> = (state = {}, action) => {
  switch (action.type) {
    case T.RECEIVED_HISTORY_PLAYBACKS:
    case T.RECEIVED_EPISODE_PLAYBACKS:
      return {
        ...state,
        ...action.playbacks.reduce<{
          [episodeId: string]: EpisodePlayback
        }>((acc, playback) => ({ ...acc, [playback.episodeId]: playback }), {}),
      }
    default:
      return state
  }
}

export default combineReducers({
  currentUserId,
  users,
  playback,
})
