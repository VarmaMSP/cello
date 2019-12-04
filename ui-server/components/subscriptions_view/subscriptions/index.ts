import { connect } from 'react-redux'
import { makeGetSubscriptions } from 'selectors/entities/podcasts'
import { AppState } from 'store'
import Subscriptions, { StateToProps } from './subscriptions'

function makeMapStateToProps() {
  const getSubscriptions = makeGetSubscriptions()

  return (state: AppState): StateToProps => ({
    subscriptions: getSubscriptions(state),
  })
}

export default connect<StateToProps, {}, {}, AppState>(makeMapStateToProps())(
  Subscriptions,
)
