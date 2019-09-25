import { getUserFeed } from 'actions/user'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import FeedPage, { DispatchToProps, OwnProps, StateToProps } from './feed'

function mapStateToProps(state: AppState): StateToProps {
  return {
    reqState: state.requests.user.getUserFeed,
  }
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    loadFeed: bindActionCreators(getUserFeed, dispatch),
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(FeedPage)
