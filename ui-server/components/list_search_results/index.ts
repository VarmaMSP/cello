import { connect } from 'react-redux'
import { makeGetSearchPodcastResults } from 'selectors/entities/search'
import { AppState } from 'store'
import ListSearchResults, {
  OwnProps,
  StateToProps,
} from './list_search_results'

function makeMapStateToProps() {
  const getSearchPodcastResults = makeGetSearchPodcastResults()

  return (state: AppState, { query }: OwnProps): StateToProps => ({
    podcasts: getSearchPodcastResults(state, query),
  })
}

export default connect<StateToProps, {}, OwnProps, AppState>(
  makeMapStateToProps(),
)(ListSearchResults)