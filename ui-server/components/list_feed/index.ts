import { Dispatch } from 'react'
import { connect } from 'react-redux'
import { makeGetUserFeed } from 'selectors/entities/episodes'
import { AppState } from 'store'
import { AppActions, SHOW_EPISODE_MODAL } from 'types/actions'
import ListFeed, { DispatchToProps, StateToProps } from './list_feed'

function makeMapStateToProps() {
  const getUserFeed = makeGetUserFeed()

  return (state: AppState): StateToProps => ({
    feed: getUserFeed(state),
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
