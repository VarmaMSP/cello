import { Dispatch, bindActionCreators } from 'redux'
import { getPodcast } from '../../actions/podcast'
import PodcastPage, { OwnProps, StateToProps, DispatchToProps } from './podcast'
import { connect } from 'react-redux'
import { AppState } from 'store'
import { makeGetEpisodesInPodcast } from '../../selectors/entities/episodes'
import { getPodcastById } from '../../selectors/entities/podcasts'
import { PLAY_EPISODE } from '../../types/actions'

function makeMapStateToProps() {
  const getEpisodesInPodcast = makeGetEpisodesInPodcast()

  return (state: AppState, props: OwnProps): StateToProps => {
    return {
      podcast: getPodcastById(state, props.id),
      episodes: getEpisodesInPodcast(state, props.id),
    }
  }
}

function mapDispatchToProps(dispatch: Dispatch): DispatchToProps {
  return {
    getPodcast: bindActionCreators(getPodcast, dispatch),
    playEpisode: (episodeId: string) =>
      dispatch({
        type: PLAY_EPISODE,
        episodeId,
      }),
  }
}

export default connect<StateToProps, DispatchToProps, OwnProps, AppState>(
  makeMapStateToProps(),
  mapDispatchToProps,
)(PodcastPage)
