import { connect } from 'react-redux'
import { Dispatch } from 'redux'
import { getEpisodeById } from 'selectors/entities/episodes'
import { getPodcastById } from 'selectors/entities/podcasts'
import { AppState } from 'store'
import * as T from 'types/actions'
import PlaylistEpisodeListItem, { DispatchToProps, OwnProps, StateToProps } from './playlist_episode_list_item'

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
    playEpisode: () =>
      dispatch({
        type: T.AUDIO_PLAYER_PLAY_EPISODE,
        episodeId: episodeId,
        beginAt: 0,
      }),
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps, mapDispatchToProps)(
  PlaylistEpisodeListItem,
)
