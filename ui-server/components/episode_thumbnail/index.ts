import { startPlayback } from 'actions/playback'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getEpisodeById } from 'selectors/entities/episodes'
import { getPodcastById } from 'selectors/entities/podcasts'
import { AppState } from 'store'
import * as T from 'types/actions'
import EpisodeThumbnail, {
  DispatchToProps,
  OwnProps,
  StateToProps,
} from './episode_thumbnail'

function mapStateToProps(
  state: AppState,
  { episodeId }: OwnProps,
): StateToProps {
  const episode = getEpisodeById(state, episodeId)
  const podcast = getPodcastById(state, episode.podcastId)

  return { episode, podcast }
}

function mapDispatchToProps(
  dispatch: Dispatch<T.AppActions>,
  { episodeId }: OwnProps,
): DispatchToProps {
  return {
    playEpisode: (beginAt: number) =>
      bindActionCreators(startPlayback, dispatch)(episodeId, beginAt),
  }
}

export default connect<StateToProps, DispatchToProps, OwnProps, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(EpisodeThumbnail)
