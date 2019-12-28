import { Dispatch } from 'react'
import { connect } from 'react-redux'
import { makeGetPlaylistsByUser } from 'selectors/entities/playlists'
import { getSignedInUserId } from 'selectors/session'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import Feed, { DispatchToProps, StateToProps } from './feed'

function makeMapStateToProps() {
  const getPlaylists = makeGetPlaylistsByUser()

  return (state: AppState): StateToProps => {
    return {
      playlists: getPlaylists(state, getSignedInUserId(state)),
      receivedAll: true,
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
