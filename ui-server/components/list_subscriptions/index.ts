import { connect } from 'react-redux'
import { makeGetUserSubscriptions } from 'selectors/entities/users'
import { AppState } from 'store'
import ListSubscriptions, { StateToProps } from './list_subscriptions'

function mapStateToProps(state: AppState): StateToProps {
  return {
    subscriptions: makeGetUserSubscriptions()(state),
  }
}

export default connect<StateToProps, {}, {}, AppState>(mapStateToProps)(
  ListSubscriptions,
)
