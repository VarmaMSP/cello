import { getSubscriptionsFeed_ } from 'actions/subscriptions'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { makeGetSubscriptionsFeed } from 'selectors/entities/feed'
import { getSubscriptionsFeedStatus } from 'selectors/request'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import Feed, { DispatchToProps, StateToProps } from './feed'

function makeMapStateToProps() {
  const getSubscriptionsFeed = makeGetSubscriptionsFeed()

  return (state: AppState): StateToProps => ({
    feed: getSubscriptionsFeed(state),
    isLoadingMore: getSubscriptionsFeedStatus(state) === 'IN_PROGRESS',
  })
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    loadMore: (offset: number) =>
      bindActionCreators(getSubscriptionsFeed_, dispatch)(offset, 20),
  }
}

export default connect<StateToProps, DispatchToProps, {}, AppState>(
  makeMapStateToProps(),
  mapDispatchToProps,
)(Feed)
