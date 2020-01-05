import { combineReducers } from 'redux'
import audioPlayer from './audio_player'
import historyFeed from './history_feed'
import modalManager from './modal_manager'
import podcastEpisodeList from './podcast_episode_list'
import searchBar from './search_bar'
import resultsList from './search_results_list'
import subscriptionsFeed from './subscriptions_feed'

export default combineReducers({
  audioPlayer,
  modalManager,
  historyFeed,
  subscriptionsFeed,
  podcastEpisodeList,
  resultsList,
  searchBar,
})
