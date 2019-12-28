import { syncPlayback } from 'actions/playback'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getViewportSize } from 'selectors/browser/viewport'
import { getEpisodeById } from 'selectors/entities/episodes'
import { getPodcastById } from 'selectors/entities/podcasts'
import {
  getCurrentTime,
  getDuration,
  getExpandOnMobile,
  getPlaybackRate,
  getPlayingEpisodeId,
  getState,
  getVolume,
} from 'selectors/ui/audio_player'
import { AppState } from 'store'
import * as T from 'types/actions'
import { AudioState } from 'types/app'
import AudioPlayer, { DispatchToProps, StateToProps } from './audio_player'

function makeMapStateToProps() {
  return (state: AppState): StateToProps => {
    const playingEpisodeId = getPlayingEpisodeId(state)
    const episode = getEpisodeById(state, playingEpisodeId)
    const podcast = !!episode && getPodcastById(state, episode.podcastId)

    return {
      episodeId: playingEpisodeId,
      episode: episode,
      podcast: podcast,
      duration: getDuration(state),
      audioState: getState(state),
      currentTime: getCurrentTime(state),
      volume: getVolume(state),
      playbackRate: getPlaybackRate(state),
      viewportSize: getViewportSize(state),
      expandOnMobile: getExpandOnMobile(state),
    }
  }
}

function mapDispatchToProps(dispatch: Dispatch<T.AppActions>): DispatchToProps {
  return {
    syncPlayback: bindActionCreators(syncPlayback, dispatch),
    setDuration: (t: number) =>
      dispatch({ type: T.AUDIO_PLAYER_UPDATE_DURATION, duration: t }),
    setAudioState: (s: AudioState) =>
      dispatch({ type: T.AUDIO_PLAYER_UPDATE_STATE, state: s }),
    setCurrentTime: (t: number) =>
      dispatch({ type: T.AUDIO_PLAYER_UPDATE_CURRENT_TIME, currentTime: t }),
    setVolume: (v: number) =>
      dispatch({ type: T.AUDIO_PLAYER_UPDATE_VOLUME, volume: v }),
    setPlaybackRate: (r: number) =>
      dispatch({ type: T.AUDIO_PLAYER_UPDATE_PLAYBACK_RATE, playbackRate: r }),
    toggleExpandOnMobile: () =>
      dispatch({ type: T.AUDIO_PLAYER_TOGGLE_EXPAND }),
  }
}

export default connect<StateToProps, DispatchToProps, {}, AppState>(
  makeMapStateToProps(),
  mapDispatchToProps,
)(AudioPlayer)
