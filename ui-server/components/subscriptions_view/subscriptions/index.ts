import { connect } from 'react-redux'
import { makeGetCurrentUserSubscriptions } from 'selectors/entities/podcasts'
import { AppState } from 'store'
import Subscriptions, { StateToProps } from './subscriptions'

function makeMapStateToProps() {
  const getCurrentUserSubscriptions = makeGetCurrentUserSubscriptions()

  return (state: AppState): StateToProps => ({
    subscriptions: getCurrentUserSubscriptions(state),
  })
}

export default connect<StateToProps, {}, {}, AppState>(makeMapStateToProps())(
  Subscriptions,
)
