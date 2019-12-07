import { connect } from 'react-redux'
import { makeGetRecommendedPodcasts } from 'selectors/entities/podcast_lists'
import { AppState } from 'store'
import Recommended, { StateToProps } from './recommended'

function makeMapStateToProps() {
  const getRecommended = makeGetRecommendedPodcasts()

  return (state: AppState): StateToProps => {
    return {
      recommended: getRecommended(state),
    }
  }
}

export default connect<StateToProps, {}, {}, AppState>(makeMapStateToProps())(
  Recommended,
)
