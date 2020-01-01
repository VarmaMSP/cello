import { getEpisodePlaybacks } from 'actions/playback'
import { loadAndShowAddToPlaylistModal } from 'actions/playlist'
import { getPodcastEpisodes as getPodcastEpisodes_ } from 'actions/podcast'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getEpisodesByIds } from 'selectors/entities/episodes'
import { getPodcastById } from 'selectors/entities/podcasts'
import { getPodcastEpisodesStatus } from 'selectors/request'
import { makeSelectPodcastEpisodeList } from 'selectors/ui/podcast_episodes_list'
import { AppState } from 'store'
import * as T from 'types/actions'
import ListEpisodes, { DispatchToProps, OwnProps, StateToProps } from './episode_list'

function makeMapStateToProps() {
  const selectPodcastEpisodeList = makeSelectPodcastEpisodeList()

  return (state: AppState, { podcastId }: OwnProps): StateToProps => {
    const [episodeIds, receivedAll] = selectPodcastEpisodeList(state, {
      podcastId,
      order: 'pub_date_desc',
    })

    return {
      podcast: getPodcastById(state, podcastId),
      episodes: getEpisodesByIds(state, episodeIds),
      receivedAll,
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
    showAddToPlaylistModal: bindActionCreators(
      loadAndShowAddToPlaylistModal,
      dispatch,
    ),
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
