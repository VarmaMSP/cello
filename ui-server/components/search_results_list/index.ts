import { getResults } from 'actions/results'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getText } from 'selectors/ui/search_bar'
import { getResultType, getSortBy, makeGetEpisodes, makeGetPodcasts } from 'selectors/ui/search_results_list'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import SearchResultsList, { DispatchToProps, StateToProps } from './search_results_list'
import { getResultsStatus } from 'selectors/request'

function makeMapStateToProps() {
  const getPodcasts = makeGetPodcasts()
  const getEpisodes = makeGetEpisodes()

  return (state: AppState): StateToProps => {
    const [podcastIds, receivedAll] = getPodcasts(state)
    const [episodeIds, receivedAll_] = getEpisodes(state)

    return {
      searchBarText: getText(state),
      resultType: getResultType(state),
      sortBy: getSortBy(state),
      podcastIds,
      episodeIds,
      receivedAll: receivedAll || receivedAll_,
      isLoadingMore: getResultsStatus(state) === 'IN_PROGRESS'
    }
  }
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    loadMore: bindActionCreators(getResults, dispatch),
  }
}

export default connect<StateToProps, DispatchToProps, {}, AppState>(
  makeMapStateToProps(),
  mapDispatchToProps,
)(SearchResultsList)
