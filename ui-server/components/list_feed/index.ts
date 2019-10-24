import { Dispatch } from 'react'
import { connect } from 'react-redux'
import { makeGetCurrentUserFeed } from 'selectors/entities/episodes'
import { AppState } from 'store'
import { AppActions, SHOW_EPISODE_MODAL } from 'types/actions'
import ListFeed, { DispatchToProps, StateToProps } from './list_feed'

function makeMapStateToProps() {
  const getCurrentUserFeed = makeGetCurrentUserFeed()

  return (state: AppState): StateToProps => ({
    feed: getCurrentUserFeed(state),
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
)(ListFeed)
