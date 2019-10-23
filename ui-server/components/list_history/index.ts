import { Dispatch } from 'react'
import { connect } from 'react-redux'
import { makeGetUserHistory } from 'selectors/entities/episodes'
import { AppState } from 'store'
import { AppActions, SHOW_EPISODE_MODAL } from 'types/actions'
import ListHistory, { DispatchToProps, StateToProps } from './list_history'

function makeMapStateToProps() {
  const getUserHistory = makeGetUserHistory()

  return (state: AppState): StateToProps => ({
    history: getUserHistory(state),
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
