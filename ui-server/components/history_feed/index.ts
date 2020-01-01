import { getHistoryFeed } from 'actions/history'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getEpisodesByIds } from 'selectors/entities/episodes'
import { getHistoryFeedStatus } from 'selectors/request'
import { getReceivedAll, makeGetEpisodeIds } from 'selectors/ui/history_feed'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import HistoryFeed, { DispatchToProps, StateToProps } from './history_feed'

function makeMapStateToProps() {
  const getEpisodeIds = makeGetEpisodeIds()

  return (state: AppState): StateToProps => {
    return {
      history: getEpisodesByIds(state, getEpisodeIds(state)),
      receivedAll: getReceivedAll(state),
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
)(HistoryFeed)
