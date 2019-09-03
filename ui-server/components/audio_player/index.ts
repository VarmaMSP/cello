import AudioPlayer, { StateToProps, DispatchToProps } from './audio_player'
import { AppState } from '../../store'
import {
  getAudioState,
  makeGetPlayingEpisode,
  makeGetPlayingPodcast,
  getPlayingEpisodeId,
  getExpandOnMobile,
} from '../../selectors/ui/player'
import { connect } from 'react-redux'
import { Dispatch } from 'redux'
import {
  AppActions,
  SET_AUDIO_STATE,
  TOGGLE_EXPAND_ON_MOBILE,
} from '../../types/actions'
import { AudioState } from '../../types/app'
import { getScreenWidth } from '../../selectors/ui/screen'

function makeMapStateToProps() {
  const getPlayingEpisode = makeGetPlayingEpisode()
  const getPlayingPodcast = makeGetPlayingPodcast()

  return (state: AppState): StateToProps => ({
    episodeId: getPlayingEpisodeId(state),
    episode: getPlayingEpisode(state),
    podcast: getPlayingPodcast(state),
    audioState: getAudioState(state),
    screenWidth: getScreenWidth(state),
    expandOnMobile: getExpandOnMobile(state),
  })
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    setAudioState: (s: AudioState) =>
      dispatch({ type: SET_AUDIO_STATE, state: s }),
    toggleExpandOnMobile: () => dispatch({ type: TOGGLE_EXPAND_ON_MOBILE }),
  }
}

export default connect<StateToProps, DispatchToProps, {}, AppState>(
  makeMapStateToProps(),
  mapDispatchToProps,
)(AudioPlayer)
