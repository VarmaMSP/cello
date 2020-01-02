import { connect } from 'react-redux'
import { makeSelectPodcasts } from 'selectors/ui/search_results_list'
import { AppState } from 'store'
import SearchResultsList, {
  OwnProps,
  StateToProps,
} from './search_results_list'

function makeMapStateToProps() {
  const selectPodcasts = makeSelectPodcasts()

  return (state: AppState, { searchQuery }: OwnProps): StateToProps => {
    const [podcastIds, receivedAll] = selectPodcasts(state, { searchQuery })

    return { podcastIds, receivedAll }
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(
  makeMapStateToProps(),
)(SearchResultsList)
