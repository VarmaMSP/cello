import {
  addEpisodeToPlaylist,
  removeEpisodeFromPlaylist,
} from 'actions/playlist'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import PlaylistsListItem, {
  DispatchToProps,
  OwnProps,
} from './playlists_list_item'

function mapDispatchToProps(
  dispatch: Dispatch<AppActions>,
  { playlist, episodeId }: OwnProps,
): DispatchToProps {
  return {
    addEpisode: () =>
      bindActionCreators(addEpisodeToPlaylist, dispatch)(
        playlist.id,
        episodeId,
      ),
    removeEpisode: () =>
      bindActionCreators(removeEpisodeFromPlaylist, dispatch)(
        playlist.id,
        episodeId,
      ),
  }
}

export default connect<{}, DispatchToProps, OwnProps, AppState>(
  null,
  mapDispatchToProps,
)(PlaylistsListItem)
