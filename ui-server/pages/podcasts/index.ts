import { Dispatch, bindActionCreators } from 'redux'
import { getPodcast } from '../../actions/podcast'
import PodcastPage from './podcast'
import { connect } from 'react-redux'
import { AppState } from 'store'
import { makeGetEpisodesInPodcast } from '../../selectors/entities/episodes'
import { getPodcastById } from '../../selectors/entities/podcasts'

function makeMapStateToProps() {
  const getEpisodesInPodcast = makeGetEpisodesInPodcast()

  return (state: AppState, props: any) => {
    return {
      podcast: getPodcastById(state, props.id),
      episodes: getEpisodesInPodcast(state, props.id),
    }
  }
}

function mapDispatchToProps(dispatch: Dispatch) {
  return {
    getPodcast: bindActionCreators(getPodcast, dispatch),
  }
}

export default connect(
  makeMapStateToProps(),
  mapDispatchToProps,
)(PodcastPage)
