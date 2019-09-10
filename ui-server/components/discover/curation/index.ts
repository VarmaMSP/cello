import { connect } from 'react-redux'
import { getCurationById } from 'selectors/entities/curations'
import { makeGetPodcastsInCuration } from 'selectors/entities/podcasts'
import { AppState } from 'store'
import CurationView, { OwnProps, StateToProps } from './curation'

function makeMapStateToProps() {
  const getPodcastsInCuration = makeGetPodcastsInCuration()

  return (state: AppState, { curationId }: OwnProps): StateToProps => ({
    curation: getCurationById(state, curationId),
    podcasts: getPodcastsInCuration(state, curationId),
  })
}

export default connect<StateToProps, {}, OwnProps, AppState>(
  makeMapStateToProps(),
)(CurationView)
