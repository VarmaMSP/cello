import { connect } from 'react-redux'
import { getPodcastById } from 'selectors/entities/podcasts'
import { getPodcastSearchResultById } from 'selectors/entities/search_results'
import { AppState } from 'store'
import ResultPodcastItem, {
  OwnProps,
  StateToProps,
} from './result_podcast_item'

function mapStateToProps(
  state: AppState,
  { podcastId, searchQuery }: OwnProps,
): StateToProps {
  return {
    podcast: getPodcastById(state, podcastId),
    podcastSearchResult: getPodcastSearchResultById(
      state,
      searchQuery,
      podcastId,
    ),
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  ResultPodcastItem,
)
