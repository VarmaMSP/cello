import { connect } from 'react-redux'
import { Dispatch } from 'redux'
import { getViewportSize } from 'selectors/browser/viewport'
import {
  getAudioState,
  getExpandOnMobile,
  getPlayingEpisodeId,
  makeGetPlayingEpisode,
  makeGetPlayingPodcast,
} from 'selectors/ui/player'
import { AppState } from 'store'
import * as T from 'types/actions'
import { AudioState } from 'types/app'
import AudioPlayer, { DispatchToProps, StateToProps } from './audio_player'

function makeMapStateToProps() {
  const getPlayingEpisode = makeGetPlayingEpisode()
  const getPlayingPodcast = makeGetPlayingPodcast()

  return (state: AppState): StateToProps => ({
    episodeId: getPlayingEpisodeId(state),
    episode: getPlayingEpisode(state),
    podcast: getPlayingPodcast(state),
    audioState: getAudioState(state),
    viewportSize: getViewportSize(state),
    expandOnMobile: getExpandOnMobile(state),
  })
}

function mapDispatchToProps(dispatch: Dispatch<T.AppActions>): DispatchToProps {
  return {
    setAudioState: (s: AudioState) =>
      dispatch({ type: T.SET_AUDIO_STATE, state: s }),
    toggleExpandOnMobile: () => dispatch({ type: T.TOGGLE_EXPAND_ON_MOBILE }),
  }
}

export default connect<StateToProps, DispatchToProps, {}, AppState>(
  makeMapStateToProps(),
  mapDispatchToProps,
)(AudioPlayer)
