import { connect } from 'react-redux'
import { Dispatch } from 'redux'
import { getPlaybackRate, getVolume } from 'selectors/ui/audio_player'
import { AppState } from 'store'
import * as T from 'types/actions'
import AudioPlayerSettings, {
  DispatchToProps,
  StateToProps,
} from './audio_player_settings'

function mapStateToProps(state: AppState): StateToProps {
  return {
    volume: getVolume(state),
    playbackRate: getPlaybackRate(state),
  }
}

function mapDispatchToProps(dispatch: Dispatch<T.AppActions>): DispatchToProps {
  return {
    changeVolume: (v: number) =>
      dispatch({ type: T.AUDIO_PLAYER_UPDATE_VOLUME, volume: v }),
    changePlaybackRate: (r: number) =>
      dispatch({ type: T.AUDIO_PLAYER_UPDATE_PLAYBACK_RATE, playbackRate: r }),
  }
}

export default connect<StateToProps, DispatchToProps, {}, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(AudioPlayerSettings)
