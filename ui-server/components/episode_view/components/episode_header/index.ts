import { connect } from 'react-redux'
import { getPodcastById } from 'selectors/entities/podcasts'
import { AppState } from 'store'
import EpisodeHeader, { OwnProps, StateToProps } from './episode_header'

function mapStateToProps(state: AppState, { episode }: OwnProps): StateToProps {
  return {
    podcast: getPodcastById(state, episode.podcastId),
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  EpisodeHeader,
)
