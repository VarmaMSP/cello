import { connect } from 'react-redux'
import { makeGetCurrentUserSubscriptions } from 'selectors/entities/podcasts'
import { AppState } from 'store'
import ListSubscriptions, { StateToProps } from './list_subscriptions'

function makeMapStateToProps() {
  const getCurrentUserSubscriptions = makeGetCurrentUserSubscriptions()

  return (state: AppState): StateToProps => ({
    subscriptions: getCurrentUserSubscriptions(state),
  })
}

export default connect<StateToProps, {}, {}, AppState>(makeMapStateToProps())(
  ListSubscriptions,
)
