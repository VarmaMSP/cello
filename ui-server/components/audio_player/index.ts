import { connect } from 'react-redux'
import { Dispatch } from 'redux'
import {
  getAudioState,
  getExpandOnMobile,
  getPlayingEpisodeId,
  makeGetPlayingEpisode,
  makeGetPlayingPodcast,
} from 'selectors/ui/player'
import { getScreenWidth } from 'selectors/ui/screen'
import { AppState } from 'store'
import {
  AppActions,
  SET_AUDIO_STATE,
  TOGGLE_EXPAND_ON_MOBILE,
} from 'types/actions'
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
