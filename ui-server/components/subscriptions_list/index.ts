import { connect } from 'react-redux'
import { makeGetUserSubscriptions } from 'selectors/entities/users'
import { AppState } from 'store'
import SubscriptionsList, { StateToProps } from './subscriptions_list'

function mapStateToProps(state: AppState): StateToProps {
  return {
    subscriptions: makeGetUserSubscriptions()(state),
  }
}

export default connect<StateToProps, {}, {}, AppState>(mapStateToProps)(
  SubscriptionsList,
)
