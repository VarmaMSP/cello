import { connect } from 'react-redux'
import { getPodcastById } from 'selectors/entities/podcasts'
import { AppState } from 'store'
import PodcastInfo, { OwnProps, StateToProps } from './podcast_info'

function mapStateToProps(state: AppState, props: OwnProps): StateToProps {
  return {
    podcast: getPodcastById(state, props.podcastId),
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  PodcastInfo,
)
