import { connect } from 'react-redux'
import { makeGetPodcastsInCuration } from 'selectors/entities/curations'
import { AppState } from 'store'
import Recommended, { StateToProps } from './recommended'

function makeMapStateToProps() {
  const getPodcastsInCuration = makeGetPodcastsInCuration()

  return (state: AppState): StateToProps => {
    return {
      podcasts: getPodcastsInCuration(state, 'recommended'),
    }
  }
}

export default connect<StateToProps, {}, {}, AppState>(makeMapStateToProps())(
  Recommended,
)
