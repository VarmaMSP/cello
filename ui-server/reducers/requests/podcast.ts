import { combineReducers } from 'redux'
import {
  GET_PODCAST_FAILURE,
  GET_PODCAST_REQUEST,
  GET_PODCAST_SUCCESS,
  GET_TRENDING_PODCASTS_FAILURE,
  GET_TRENDING_PODCASTS_REQUEST,
  GET_TRENDING_PODCASTS_SUCCESS,
} from 'types/actions'
import { defaultRequestReducer } from './utils'

const getPodcast = defaultRequestReducer(
  GET_PODCAST_REQUEST,
  GET_PODCAST_SUCCESS,
  GET_PODCAST_FAILURE,
)

const getTrendingPodcasts = defaultRequestReducer(
  GET_TRENDING_PODCASTS_REQUEST,
  GET_TRENDING_PODCASTS_SUCCESS,
  GET_TRENDING_PODCASTS_FAILURE,
)

export default combineReducers({
  getPodcast,
  getTrendingPodcasts,
})
