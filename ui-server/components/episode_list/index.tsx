import { connect } from 'react-redux'
import { Dispatch } from 'redux'
import { makeGetEpisodesInPodcast } from 'selectors/entities/episodes'
import { AppState } from 'store'
import { AppActions, PLAY_EPISODE } from 'types/actions'
import EpisodeList, {
  DispatchToProps,
  OwnProps,
  StateToProps,
} from './episode_list'

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

export default connect<StateToProps, DispatchToProps, OwnProps, AppState>(
  makeMapStateToProps(),
  dispatchToProps,
)(EpisodeList)
