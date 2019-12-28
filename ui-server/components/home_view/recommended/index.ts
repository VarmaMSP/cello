import { connect } from 'react-redux'
import { makeGetPodcastsInCuration } from 'selectors/entities/curations'
import { getPodcastsByIds } from 'selectors/entities/podcasts'
import { AppState } from 'store'
import Recommended, { StateToProps } from './recommended'

function makeMapStateToProps() {
  const getPodcastsInCuration = makeGetPodcastsInCuration()

  return (state: AppState): StateToProps => {
    return {
      podcasts: getPodcastsByIds(
        state,
        getPodcastsInCuration(state, 'recommended'),
      ),
    }
  }
}

export default connect<StateToProps, {}, {}, AppState>(makeMapStateToProps())(
  Recommended,
)
