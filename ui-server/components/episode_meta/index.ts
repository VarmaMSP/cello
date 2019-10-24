import { connect } from 'react-redux'
import { getCurrentUserPlayback } from 'selectors/entities/episodes'
import { AppState } from 'store'
import EpisodeMeta, { OwnProps, StateToProps } from './episode_meta'

function mapStateToProps(state: AppState, { episode }: OwnProps): StateToProps {
  return {
    playback: getCurrentUserPlayback(state, episode.id),
  }
}

export default connect<StateToProps, null, OwnProps, AppState>(mapStateToProps)(
  EpisodeMeta,
)
