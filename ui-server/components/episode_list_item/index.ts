import { beginPlayback } from 'actions/episode'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getCurrentUserPlayback, getEpisodeById } from 'selectors/entities/episodes'
import { getPodcastById } from 'selectors/entities/podcasts'
import { AppState } from 'store'
import * as T from 'types/actions'
import { AppActions } from 'types/actions'
import EpisodeListItem, { DispatchToProps, OwnProps, StateToProps } from './episode_list_item'

function mapStateToProps(
  state: AppState,
  { episodeId }: OwnProps,
): StateToProps {
  const episode = getEpisodeById(state, episodeId)
  const podcast = getPodcastById(state, episode.podcastId)
  return {
    episode,
    podcast,
    playback: getCurrentUserPlayback(state, episodeId),
  }
}

function mapDispatchToProps(
  dispatch: Dispatch<AppActions>,
  { episodeId }: OwnProps,
): DispatchToProps {
  return {
    playEpisode: (startTime: number) =>
      bindActionCreators(beginPlayback, dispatch)(episodeId, startTime),
    showAddToPlaylistModal: () =>
      dispatch({
        type: T.SHOW_ADD_TO_PLAYLIST_MODAL,
        episodeId,
      }),
  }
}

export default connect<StateToProps, DispatchToProps, OwnProps, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(EpisodeListItem)