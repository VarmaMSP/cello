import { Dispatch } from 'react'
import { connect } from 'react-redux'
import { makeGetUserFeed } from 'selectors/entities/episodes'
import { AppState } from 'store'
import { AppActions, PLAY_EPISODE } from 'types/actions'
import ListFeed, { DispatchToProps, StateToProps } from './list_feed'

function makeMapStateToProps() {
  const getUserFeed = makeGetUserFeed()

  return (state: AppState): StateToProps => ({
    feed: getUserFeed(state),
  })
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>) {
  return {
    playEpisode: (episodeId: string) =>
      dispatch({
        type: PLAY_EPISODE,
        episodeId,
      }),
  }
}

export default connect<StateToProps, DispatchToProps, {}, AppState>(
  makeMapStateToProps(),
  mapDispatchToProps,
)(ListFeed)
