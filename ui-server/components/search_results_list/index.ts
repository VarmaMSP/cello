import { connect } from 'react-redux'
import { makeSelectEpisodes, makeSelectPodcasts } from 'selectors/ui/search_results_list'
import { AppState } from 'store'
import SearchResultsList, { OwnProps, StateToProps } from './search_results_list'

function makeMapStateToProps() {
  const selectPodcasts = makeSelectPodcasts()
  const selectEpisodes = makeSelectEpisodes()

  return (state: AppState, { searchQuery, sortBy }: OwnProps): StateToProps => {
    const [podcastIds, receivedAll] = selectPodcasts(state, {
      searchQuery,
      sortBy,
    })
    const [episodeIds, receivedAll_] = selectEpisodes(state, {
      searchQuery,
      sortBy,
    })

    return { podcastIds, episodeIds, receivedAll: receivedAll || receivedAll_ }
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(
  makeMapStateToProps(),
)(SearchResultsList)
