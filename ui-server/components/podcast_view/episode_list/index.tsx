import { getEpisodePlaybacks } from 'actions/episode'
import { getPodcastEpisodes } from 'actions/podcast'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import {
  makeGetEpisodesInPodcast,
  makeGetReceivedAllEpisodes,
} from 'selectors/entities/episodes'
import { getPodcastEpisodesStatus } from 'selectors/request'
import { AppState } from 'store'
import * as T from 'types/actions'
import ListEpisodes, {
  DispatchToProps,
  OwnProps,
  StateToProps,
} from './episode_list'

function makeMapStateToProps() {
  const getEpisodesInPodcast = makeGetEpisodesInPodcast()
  const getReceivedAllEpisodes = makeGetReceivedAllEpisodes()

  return (state: AppState, { podcastId }: OwnProps): StateToProps => ({
    episodes: getEpisodesInPodcast(state, podcastId),
    receivedAllEpisodes: getReceivedAllEpisodes(state, podcastId),
    isLoadingMore: getPodcastEpisodesStatus(state, podcastId) === 'IN_PROGRESS',
  })
}

function dispatchToProps(
  dispatch: Dispatch<T.AppActions>,
  { podcastId }: OwnProps,
): DispatchToProps {
  return {
    showAddToPlaylistModal: (episodeId: string) =>
      dispatch({
        type: T.SHOW_ADD_TO_PLAYLIST_MODAL,
        episodeId,
      }),
    loadEpisodes: (offset: number) =>
      bindActionCreators(getPodcastEpisodes, dispatch)(
        podcastId,
        20,
        offset,
        'pub_date_desc',
      ),
    loadEpisodePlaybacks: bindActionCreators(getEpisodePlaybacks, dispatch),
  }
}

export default connect<StateToProps, DispatchToProps, OwnProps, AppState>(
  makeMapStateToProps(),
  dispatchToProps,
)(ListEpisodes)
