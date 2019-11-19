import { getSubscriptionsFeed } from 'actions/subscriptions'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { makeGetCurrentUserFeed } from 'selectors/entities/episodes'
import { getSubscriptionsFeedStatus } from 'selectors/request'
import { AppState } from 'store'
import { AppActions, SHOW_EPISODE_MODAL } from 'types/actions'
import Feed, { DispatchToProps, StateToProps } from './feed'

function makeMapStateToProps() {
  const getCurrentUserFeed = makeGetCurrentUserFeed()

  return (state: AppState): StateToProps => ({
    feed: getCurrentUserFeed(state),
    isLoadingMore: getSubscriptionsFeedStatus(state) === 'IN_PROGRESS',
  })
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    loadMore: bindActionCreators(getSubscriptionsFeed, dispatch),
    showEpisodeModal: (episodeId: string) =>
      dispatch({
        type: SHOW_EPISODE_MODAL,
        episodeId,
      }),
  }
}

export default connect<StateToProps, DispatchToProps, {}, AppState>(
  makeMapStateToProps(),
  mapDispatchToProps,
)(Feed)
