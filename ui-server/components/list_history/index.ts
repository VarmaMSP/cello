import { Dispatch } from 'react'
import { connect } from 'react-redux'
import { makeGetCurrentUserHistory } from 'selectors/entities/episodes'
import { AppState } from 'store'
import { AppActions, SHOW_EPISODE_MODAL } from 'types/actions'
import ListHistory, { DispatchToProps, StateToProps } from './list_history'

function makeMapStateToProps() {
  const getCurrentUserHistory = makeGetCurrentUserHistory()

  return (state: AppState): StateToProps => ({
    history: getCurrentUserHistory(state),
  })
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
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
)(ListHistory)
