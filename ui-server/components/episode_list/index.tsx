import { AppState } from 'store'
import EpisodeList, { StateToProps, OwnProps } from './episode_list'
import { makeGetEpisodesInPodcast } from '../../selectors/entities/episodes'
import { Dispatch } from 'redux'
import { PLAY_EPISODE, AppActions } from '../../types/actions'
import { connect } from 'react-redux'

function makeMapStateToProps() {
  const getEpisodesInPodcast = makeGetEpisodesInPodcast()

  return (state: AppState, props: OwnProps): StateToProps => ({
    ...props,
    episodes: getEpisodesInPodcast(state, props.podcastId),
  })
}

function dispatchToProps(dispatch: Dispatch<AppActions>) {
  return {
    playEpisode: (episodeId: string) =>
      dispatch({
        type: PLAY_EPISODE,
        episodeId,
      }),
  }
}

export default connect(
  makeMapStateToProps(),
  dispatchToProps,
)(EpisodeList)
