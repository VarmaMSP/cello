import { getHistoryFeed } from 'actions/history'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { makeGetHistoryFeed } from 'selectors/entities/feed'
import { getHistoryFeedStatus } from 'selectors/request'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import Feed, { DispatchToProps, StateToProps } from './feed'

function makeMapStateToProps() {
  const getHistoryFeed = makeGetHistoryFeed()

  return (state: AppState): StateToProps => {
    const { episodes, receivedAll } = getHistoryFeed(state)

    return {
      history: episodes,
      receivedAll: receivedAll,
      isLoadingMore: getHistoryFeedStatus(state) === 'IN_PROGRESS',
    }
  }
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    loadMore: (offset: number) =>
      bindActionCreators(getHistoryFeed, dispatch)(offset, 20),
  }
}

export default connect<StateToProps, DispatchToProps, {}, AppState>(
  makeMapStateToProps(),
  mapDispatchToProps,
)(Feed)
