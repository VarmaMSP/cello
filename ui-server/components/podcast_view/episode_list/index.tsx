import { getEpisodePlaybacks } from 'actions/playback'
import { getPodcastEpisodes as getPodcastEpisodes_ } from 'actions/podcast'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { makeGetEpisodesByPodcast } from 'selectors/entities/episodes'
import { getPodcastById } from 'selectors/entities/podcasts'
import { getPodcastEpisodesStatus } from 'selectors/request'
import { AppState } from 'store'
import * as T from 'types/actions'
import ListEpisodes, { DispatchToProps, OwnProps, StateToProps } from './episode_list'

function makeMapStateToProps() {
  const getPodcastEpisodes = makeGetEpisodesByPodcast()

  return (state: AppState, { podcastId }: OwnProps): StateToProps => {
    return {
      podcast: getPodcastById(state, podcastId),
      episodes: getPodcastEpisodes(state, podcastId),
      receivedAll: false,
      isLoadingMore:
        getPodcastEpisodesStatus(state, podcastId) === 'IN_PROGRESS',
    }
  }
}

function dispatchToProps(
  dispatch: Dispatch<T.AppActions>,
  { podcastId }: OwnProps,
): DispatchToProps {
  return {
    showAddToPlaylistModal: (episodeId: string) =>
      dispatch({
        type: T.MODAL_MANAGER_SHOW_ADD_TO_PLAYLIST_MODAL,
        episodeId,
      }),
    loadEpisodes: (offset: number) =>
      bindActionCreators(getPodcastEpisodes_, dispatch)(
        podcastId,
        20,
        offset,
        'pub_date_desc',
      ),
    loadPlaybacks: bindActionCreators(getEpisodePlaybacks, dispatch),
  }
}

export default connect<StateToProps, DispatchToProps, OwnProps, AppState>(
  makeMapStateToProps(),
  dispatchToProps,
)(ListEpisodes)
