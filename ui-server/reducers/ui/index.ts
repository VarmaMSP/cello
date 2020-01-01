import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import audioPlayer from './audio_player'
import historyFeed from './history_feed'
import modalManager from './modal_manager'
import subscriptionsFeed from './subscriptions_feed'
import podcastEpisodeList from './podcast_episode_list'

const searchText: Reducer<string, T.AppActions> = (state = '', action) => {
  switch (action.type) {
    case T.SEARCH_BAR_TEXT_CHANGE:
      return action.text
    default:
      return state
  }
}

export default combineReducers({
  searchText,
  audioPlayer,
  modalManager,
  historyFeed,
  subscriptionsFeed,
  podcastEpisodeList,
})
