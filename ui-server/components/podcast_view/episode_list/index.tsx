import { getPodcastEpisodes as getPodcastEpisodes_ } from 'actions/episode'
import { getEpisodePlaybacks } from 'actions/playback'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { makeGetPodcastEpisodes } from 'selectors/entities/episodes'
import { getPodcastEpisodesStatus } from 'selectors/request'
import { AppState } from 'store'
import * as T from 'types/actions'
import ListEpisodes, {
  DispatchToProps,
  OwnProps,
  StateToProps,
} from './episode_list'

function makeMapStateToProps() {
  const getPodcastEpisodes = makeGetPodcastEpisodes()

  return (state: AppState, { podcastId }: OwnProps): StateToProps => {
    const { episodes, receivedAll } = getPodcastEpisodes(state, podcastId)

    return {
      episodes,
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
    showAddToPlaylistModal: (episodeId: string) =>
      dispatch({
        type: T.SHOW_ADD_TO_PLAYLIST_MODAL,
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
