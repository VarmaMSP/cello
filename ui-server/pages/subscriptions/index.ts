import { connect } from 'react-redux'
import { AppState } from 'store'
import SubscriptionsPage, { OwnProps, StateToProps } from './subscriptions'

function mapStateToProps(state: AppState): StateToProps {
  return {
    reqState: state.requests.user.getSignedInUser,
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  SubscriptionsPage,
)
