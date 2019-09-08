import { connect } from 'react-redux'
import { getPodcastById } from 'selectors/entities/podcasts'
import { AppState } from 'store'
import PodcastDetails, { OwnProps, StateToProps } from './podcast_details'

function mapStateToProps(state: AppState, props: OwnProps) {
  return {
    ...props,
    podcast: getPodcastById(state, props.podcastId),
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  PodcastDetails,
)
