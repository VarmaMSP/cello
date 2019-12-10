import { startPlayback } from 'actions/playback'
import { loadAndShowAddToPlaylistModal } from 'actions/playlist'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getEpisodeById } from 'selectors/entities/episodes'
import { getPodcastById } from 'selectors/entities/podcasts'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import EpisodeListItem, {
  DispatchToProps,
  OwnProps,
  StateToProps,
} from './episode_list_item'

function mapStateToProps(
  state: AppState,
  { episodeId }: OwnProps,
): StateToProps {
  const episode = getEpisodeById(state, episodeId)
  const podcast = getPodcastById(state, episode.podcastId)

  return { episode, podcast }
}

function mapDispatchToProps(
  dispatch: Dispatch<AppActions>,
  { episodeId }: OwnProps,
): DispatchToProps {
  return {
    playEpisode: (startTime: number) =>
      bindActionCreators(startPlayback, dispatch)(episodeId, startTime),
    showAddToPlaylistModal: () =>
      bindActionCreators(loadAndShowAddToPlaylistModal, dispatch)(episodeId),
  }
}

export default connect<StateToProps, DispatchToProps, OwnProps, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(EpisodeListItem)
