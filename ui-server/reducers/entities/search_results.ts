import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { EpisodeSearchResult, PodcastSearchResult } from 'models'

const byPodcastId: Reducer<
  { [searchQuery: string]: { [podcastId: string]: PodcastSearchResult } },
  T.AppActions
> = (state = {}, action) => {
  switch (action.type) {
    case T.SEARCH_RESULT_ADD_PODCAST:
      return {
        ...state,
        [action.searchQuery]: action.podcastSearchResults.reduce<{
          [podcastId: string]: PodcastSearchResult
        }>(
          (acc, p) => ({ ...acc, [p.id]: p }),
          state[action.searchQuery] || {},
        ),
      }

    default:
      return state
  }
}

const byEpisodeId: Reducer<
  { [searchQuery: string]: { [episodeId: string]: EpisodeSearchResult } },
  T.AppActions
> = (state = {}, action) => {
  switch (action.type) {
    case T.SEARCH_RESULT_ADD_EPISODE:
      return {
        ...state,
        [action.searchQuery]: action.episodeSearchResults.reduce<{
          [episodeId: string]: EpisodeSearchResult
        }>(
          (acc, e) => ({ ...acc, [e.id]: e }),
          state[action.searchQuery] || {},
        ),
      }

    default:
      return state
  }
}

export default combineReducers({
  byPodcastId,
  byEpisodeId,
})
