import { Playlist } from 'types/app'

export interface StateToProps {
  playlist: Playlist
  containsEpisode: boolean
}

export interface DispatchToProps {
  addEpisode: () => void
  removeEpisode: () => void
}

export interface OwnProps {
  episodeId: string
  playlistId: string
}

type Props = StateToProps & DispatchToProps & OwnProps

const AddToPlaylistItem: React.FC<Props> = ({
  playlist,
  containsEpisode,
  addEpisode,
  removeEpisode,
}) => {
  return (
    <div>
      <label className="inline-flex items-center">
        <input
          type="checkbox"
          className="form-checkbox"
          checked={containsEpisode}
          onChange={(e) => (e.target.checked ? addEpisode() : removeEpisode())}
        />
        <span className="ml-2">{playlist.title}</span>
      </label>
    </div>
  )
}

export default AddToPlaylistItem
