import { getSubscriptionsFeed } from 'actions/subscription'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { makeGetSubscriptionsFeed } from 'selectors/entities/feed'
import { getSubscriptionsFeedStatus } from 'selectors/request'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import Feed, { DispatchToProps, StateToProps } from './feed'

function makeMapStateToProps() {
  const getSubscriptionsFeed = makeGetSubscriptionsFeed()

  return (state: AppState): StateToProps => {
    const { episodes, receivedAll } = getSubscriptionsFeed(state)
    return {
      feed: episodes,
      receivedAll: receivedAll,
      isLoadingMore: getSubscriptionsFeedStatus(state) === 'IN_PROGRESS',
    }
  } 
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    loadMore: (offset: number) =>
      bindActionCreators(getSubscriptionsFeed, dispatch)(offset, 20),
  }
}

export default connect<StateToProps, DispatchToProps, {}, AppState>(
  makeMapStateToProps(),
  mapDispatchToProps,
)(Feed)
