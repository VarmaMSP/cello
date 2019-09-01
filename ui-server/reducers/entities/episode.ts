import { Episode } from 'types/app'
import { AppActions, RECEIVED_EPISODE } from 'types/actions'
import { combineReducers } from 'redux'

type EpisodesState = Readonly<{ [episodeId: string]: Episode }>

export const episodes = (
  state: EpisodesState = {},
  action: AppActions,
): EpisodesState => {
  switch (action.type) {
    case RECEIVED_EPISODE:
      return { ...state, [action.episode.id]: action.episode }
    default:
      return state
  }
}

type EpisodeByPodcastState = Readonly<{ [podcastId: string]: string[] }>

export const episodesByPodcast = (
  state: EpisodeByPodcastState = {},
  action: AppActions,
): EpisodeByPodcastState => {
  switch (action.type) {
    case RECEIVED_EPISODE:
      let episodeId = action.episode.id
      let podcastId = action.episode.podcastId
      let episodes = state[podcastId] || []
      return { ...state, [podcastId]: [...episodes, episodeId] }
    default:
      return state
  }
}

export default combineReducers({
  episodes,
  episodesByPodcast,
})
