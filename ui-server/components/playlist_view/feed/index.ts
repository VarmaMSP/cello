import { Dispatch } from 'react'
import { connect } from 'react-redux'
import { makeGetUserPlaylists } from 'selectors/entities/playlists'
import { getCurrentUserId } from 'selectors/entities/users'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import Feed, { DispatchToProps, StateToProps } from './feed'

function makeMapStateToProps() {
  const getPlaylists = makeGetUserPlaylists()

  return (state: AppState): StateToProps => {
    const { playlists, receivedAll } = getPlaylists(
      state,
      getCurrentUserId(state),
    )
    return {
      playlists,
      receivedAll,
      isLoadingMore: false,
    }
  }
}

function mapDispatchToProps(_: Dispatch<AppActions>): DispatchToProps {
  return {
    loadMore: (_: number) => {},
  }
}

export default connect<StateToProps, DispatchToProps, {}, AppState>(
  makeMapStateToProps(),
  mapDispatchToProps,
)(Feed)
