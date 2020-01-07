import { connect } from 'react-redux'
import { getEpisodeById } from 'selectors/entities/episodes'
import { getPodcastById } from 'selectors/entities/podcasts'
import { makeGetEpisodeSearchResultById } from 'selectors/entities/search_results'
import { AppState } from 'store'
import ResultEpisodeItem, {
  OwnProps,
  StateToProps,
} from './result_episode_item'

function makeMapStateToProps() {
  const getEpisodeSearchResultById = makeGetEpisodeSearchResultById()

  return (state: AppState, { episodeId }: OwnProps): StateToProps => {
    const episode = getEpisodeById(state, episodeId)
    const podcast = getPodcastById(state, episode.podcastId)
    const episodeSearchResult = getEpisodeSearchResultById(state, episodeId)

    return {
      episode,
      podcast,
      episodeSearchResult,
    }
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(
  makeMapStateToProps,
)(ResultEpisodeItem)
