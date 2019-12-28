import { connect } from 'react-redux'
import { getPodcastsByIds } from 'selectors/entities/podcasts'
import { getSubscriptions } from 'selectors/session'
import { AppState } from 'store'
import Subscriptions, { StateToProps } from './subscriptions'

function mapStateToProps(state: AppState): StateToProps {
  return {
    subscriptions: getPodcastsByIds(state, getSubscriptions(state)),
  }
}

export default connect<StateToProps, {}, {}, AppState>(mapStateToProps)(
  Subscriptions,
)
