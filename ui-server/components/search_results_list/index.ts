import { getEpisodePlaybacks } from 'actions/playback'
import { getResults } from 'actions/results'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getResultsStatus } from 'selectors/request'
import { getIsUserSignedIn } from 'selectors/session'
import { getText } from 'selectors/ui/search_bar'
import {
  getResultType,
  getSortBy,
  makeGetEpisodes,
  makeGetPodcasts,
  getQuery,
} from 'selectors/ui/search_results_list'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import SearchResultsList, {
  DispatchToProps,
  StateToProps,
} from './search_results_list'

function makeMapStateToProps() {
  const getPodcasts = makeGetPodcasts()
  const getEpisodes = makeGetEpisodes()

  return (state: AppState): StateToProps => {
    const [podcastIds, receivedAll] = getPodcasts(state)
    const [episodeIds, receivedAll_] = getEpisodes(state)
    const resultType = getResultType(state)

    return {
      isUserSignedIn: getIsUserSignedIn(state),
      searchBarText: getText(state),
      query: getQuery(state),
      resultType,
      sortBy: getSortBy(state),
      podcastIds,
      episodeIds,
      receivedAll: resultType === 'podcast' ? receivedAll : receivedAll_,
      isLoadingMore: getResultsStatus(state) === 'IN_PROGRESS',
    }
  }
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    loadMore: bindActionCreators(getResults, dispatch),
    loadPlaybacks: bindActionCreators(getEpisodePlaybacks, dispatch),
  }
}

export default connect<StateToProps, DispatchToProps, {}, AppState>(
  makeMapStateToProps(),
  mapDispatchToProps,
)(SearchResultsList)
