import { getSubscriptionsFeed } from 'actions/subscription'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getEpisodesByIds } from 'selectors/entities/episodes'
import { getSubscriptionsFeedStatus } from 'selectors/request'
import {
  getReceivedAll,
  makeGetEpisodeIds,
} from 'selectors/ui/subscriptions_feed'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import Feed, { DispatchToProps, StateToProps } from './feed'

function makeMapStateToProps() {
  const getSubscriptionsFeed = makeGetEpisodeIds()

  return (state: AppState): StateToProps => {
    return {
      feed: getEpisodesByIds(state, getSubscriptionsFeed(state)),
      receivedAll: getReceivedAll(state),
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
