import { getSubscriptionsFeed } from 'actions/subscription'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getEpisodesByIds } from 'selectors/entities/episodes'
import { getSubscriptionsFeedStatus } from 'selectors/request'
import { makeSelectSubscriptionsFeed } from 'selectors/ui/subscriptions_feed'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import SubscriptionsFeed, { DispatchToProps, StateToProps } from './subscriptions_feed'

function makeMapStateToProps() {
  const selectSubscriptionsFeed = makeSelectSubscriptionsFeed()

  return (state: AppState): StateToProps => {
    const [episodeIds, receivedAll] = selectSubscriptionsFeed(state)

    return {
      feed: getEpisodesByIds(state, episodeIds),
      receivedAll,
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
)(SubscriptionsFeed)
