import { getEpisodePlaybacks } from 'actions/episode'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { makeGetEpisodesInPodcast } from 'selectors/entities/episodes'
import { AppState } from 'store'
import * as T from 'types/actions'
import ListEpisodes, { DispatchToProps, OwnProps, StateToProps } from './episode_list'

function makeMapStateToProps() {
  const getEpisodesInPodcast = makeGetEpisodesInPodcast()

  return (state: AppState, { podcastId }: OwnProps): StateToProps => ({
    episodes: getEpisodesInPodcast(state, podcastId),
  })
}

function dispatchToProps(dispatch: Dispatch<T.AppActions>): DispatchToProps {
  return {
    showAddToPlaylistModal: (episodeId: string) =>
      dispatch({
        type: T.SHOW_ADD_TO_PLAYLIST_MODAL,
        episodeId,
      }),
    loadEpisodePlaybacks: bindActionCreators(getEpisodePlaybacks, dispatch),
  }
}

export default connect<StateToProps, DispatchToProps, OwnProps, AppState>(
  makeMapStateToProps(),
  dispatchToProps,
)(ListEpisodes)
