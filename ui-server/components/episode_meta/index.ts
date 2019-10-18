import { connect } from 'react-redux'
import { getEpisodePlayback } from 'selectors/entities/users'
import { AppState } from 'store'
import EpisodeMeta, { OwnProps, StateToProps } from './episode_meta'

function mapStateToProps(state: AppState, { episode }: OwnProps): StateToProps {
  return {
    playback: getEpisodePlayback(state, episode.id),
  }
}

export default connect<StateToProps, null, OwnProps, AppState>(mapStateToProps)(
  EpisodeMeta,
)
