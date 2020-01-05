import { connect } from 'react-redux'
import { getEpisodeById } from 'selectors/entities/episodes'
import { getPodcastById } from 'selectors/entities/podcasts'
import { getEpisodeSearchResultById } from 'selectors/entities/search_results'
import { AppState } from 'store'
import ResultEpisodeItem, {
  OwnProps,
  StateToProps,
} from './result_episode_item'

function mapStateToProps(
  state: AppState,
  { episodeId, searchQuery }: OwnProps,
): StateToProps {
  const episode = getEpisodeById(state, episodeId)
  const podcast = getPodcastById(state, episode.podcastId)

  return {
    episode,
    podcast,
    episodeSearchResult: getEpisodeSearchResultById(
      state,
      searchQuery,
      episodeId,
    ),
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  ResultEpisodeItem,
)
