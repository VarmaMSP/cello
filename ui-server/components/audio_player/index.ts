import { syncPlayback } from 'actions/episode'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getViewportSize } from 'selectors/browser/viewport'
import {
  getAudioCurrentTime,
  getAudioDuration,
  getAudioPlaybackRate,
  getAudioState,
  getAudioVolume,
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
    duration: getAudioDuration(state),
    audioState: getAudioState(state),
    currentTime: getAudioCurrentTime(state),
    volume: getAudioVolume(state),
    playbackRate: getAudioPlaybackRate(state),
    viewportSize: getViewportSize(state),
    expandOnMobile: getExpandOnMobile(state),
  })
}

function mapDispatchToProps(dispatch: Dispatch<T.AppActions>): DispatchToProps {
  return {
    syncPlayback: bindActionCreators(syncPlayback, dispatch),
    setDuration: (t: number) => dispatch({ type: T.SET_DURATION, duration: t }),
    setAudioState: (s: AudioState) =>
      dispatch({ type: T.SET_AUDIO_STATE, state: s }),
    setCurrentTime: (t: number) =>
      dispatch({ type: T.SET_CURRENT_TIME, currentTime: t }),
    setVolume: (v: number) => dispatch({ type: T.SET_VOLUME, volume: v }),
    setPlaybackRate: (r: number) =>
      dispatch({ type: T.SET_PLAYBACK_RATE, playbackRate: r }),
    toggleExpandOnMobile: () => dispatch({ type: T.TOGGLE_EXPAND_ON_MOBILE }),
  }
}

export default connect<StateToProps, DispatchToProps, {}, AppState>(
  makeMapStateToProps(),
  mapDispatchToProps,
)(AudioPlayer)
