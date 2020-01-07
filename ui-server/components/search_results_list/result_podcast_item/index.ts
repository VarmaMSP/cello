import { connect } from 'react-redux'
import { getPodcastById } from 'selectors/entities/podcasts'
import { makeGetPodcastSearchResultById } from 'selectors/entities/search_results'
import { AppState } from 'store'
import ResultPodcastItem, {
  OwnProps,
  StateToProps,
} from './result_podcast_item'

function makeMapStateToProps() {
  const getPodcastSearchResultById = makeGetPodcastSearchResultById()

  return (state: AppState, { podcastId }: OwnProps): StateToProps => {
    return {
      podcast: getPodcastById(state, podcastId),
      podcastSearchResult: getPodcastSearchResultById(state, podcastId),
    }
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(
  makeMapStateToProps(),
)(ResultPodcastItem)
