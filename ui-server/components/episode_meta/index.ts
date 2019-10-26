import { connect } from 'react-redux'
import {
  getCurrentUserPlayback,
  getEpisodeById,
} from 'selectors/entities/episodes'
import { AppState } from 'store'
import EpisodeMeta, { OwnProps, StateToProps } from './episode_meta'

function mapStateToProps(
  state: AppState,
  { episodeId }: OwnProps,
): StateToProps {
  return {
    episode: getEpisodeById(state, episodeId),
    playback: getCurrentUserPlayback(state, episodeId),
  }
}

export default connect<StateToProps, null, OwnProps, AppState>(mapStateToProps)(
  EpisodeMeta,
)
