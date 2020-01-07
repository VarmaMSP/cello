import { connect } from 'react-redux'
import {
  getResultType,
  makeGetEpisodes,
  makeGetPodcasts,
} from 'selectors/ui/search_results_list'
import { AppState } from 'store'
import SearchResultsList, { StateToProps } from './search_results_list'

function makeMapStateToProps() {
  const getPodcasts = makeGetPodcasts()
  const getEpisodes = makeGetEpisodes()

  return (state: AppState): StateToProps => {
    const [podcastIds, receivedAll] = getPodcasts(state)
    const [episodeIds, receivedAll_] = getEpisodes(state)

    return {
      resultType: getResultType(state),
      podcastIds,
      episodeIds,
      receivedAll: receivedAll || receivedAll_,
    }
  }
}

export default connect<StateToProps, {}, {}, AppState>(makeMapStateToProps())(
  SearchResultsList,
)
