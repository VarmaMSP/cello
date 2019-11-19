import { combineReducers } from 'redux'
import * as T from 'types/actions'
import { defaultRequestReducer } from './utils'

const getPodcast = defaultRequestReducer(
  T.GET_PODCAST_REQUEST,
  T.GET_PODCAST_SUCCESS,
  T.GET_PODCAST_FAILURE,
)

const getPodcastEpisodes = defaultRequestReducer(
  T.GET_PODCAST_EPISODES_REQUEST,
  T.GET_PODCAST_EPISODES_SUCCESS,
  T.GET_PODCAST_EPISODES_FAILURE,
)

const getTrendingPodcasts = defaultRequestReducer(
  T.GET_TRENDING_PODCASTS_REQUEST,
  T.GET_TRENDING_PODCASTS_SUCCESS,
  T.GET_TRENDING_PODCASTS_FAILURE,
)

export default combineReducers({
  getPodcast,
  getPodcastEpisodes,
  getTrendingPodcasts,
})
