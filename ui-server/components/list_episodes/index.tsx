import { connect } from 'react-redux'
import { Dispatch } from 'redux'
import { makeGetEpisodesInPodcast } from 'selectors/entities/episodes'
import { AppState } from 'store'
import * as T from 'types/actions'
import ListEpisodes, {
  DispatchToProps,
  OwnProps,
  StateToProps,
} from './list_episodes'

function makeMapStateToProps() {
  const getEpisodesInPodcast = makeGetEpisodesInPodcast()

  return (state: AppState, props: OwnProps): StateToProps => ({
    ...props,
    episodes: getEpisodesInPodcast(state, props.podcastId),
  })
}

function dispatchToProps(dispatch: Dispatch<T.AppActions>) {
  return {
    playEpisode: (episodeId: string) =>
      dispatch({
        type: T.PLAY_EPISODE,
        episodeId,
      }),
  }
}

export default connect<StateToProps, DispatchToProps, OwnProps, AppState>(
  makeMapStateToProps(),
  dispatchToProps,
)(ListEpisodes)
