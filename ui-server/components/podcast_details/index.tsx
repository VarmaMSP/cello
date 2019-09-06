import PodcastDetails, { OwnProps, StateToProps } from './podcast_details'
import { AppState } from '../../store'
import { getPodcastById } from '../../selectors/entities/podcasts'
import { connect } from 'react-redux'

function mapStateToProps(state: AppState, props: OwnProps) {
  return {
    ...props,
    podcast: getPodcastById(state, props.podcastId),
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  PodcastDetails,
)
