import { connect } from 'react-redux'
import { getPodcastById } from 'selectors/entities/podcasts'
import { AppState } from 'store'
import HomeTab, { OwnProps, StateToProps } from './home_tab'

function mapStateToProps(state: AppState, props: OwnProps): StateToProps {
  return {
    podcast: getPodcastById(state, props.podcastId),
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  HomeTab,
)
