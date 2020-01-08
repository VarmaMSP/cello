import { connect } from 'react-redux'
import { getPodcastById } from 'selectors/entities/podcasts'
import { makeGetPodcastSearchResultById } from 'selectors/entities/search_results'
import { AppState } from 'store'
import PodcastPreview, { OwnProps, StateToProps } from './podcast_preview'

function makeMapStateToProps() {
  const getPodcastSearchResultById = makeGetPodcastSearchResultById()

  return (state: AppState, { podcastId }: OwnProps): StateToProps => {
    const podcast = getPodcastById(state, podcastId)
    const podcastSearchResult = getPodcastSearchResultById(state, podcastId)

    return { podcast, podcastSearchResult }
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(
  makeMapStateToProps(),
)(PodcastPreview)
