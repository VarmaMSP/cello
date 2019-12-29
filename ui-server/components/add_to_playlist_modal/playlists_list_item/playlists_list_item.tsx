import { Playlist } from 'types/app'

export interface DispatchToProps {
  addEpisode: () => void
  removeEpisode: () => void
}

export interface OwnProps {
  playlist: Playlist
  episodeId: string
}

const PlaylistsListItem: React.FC<DispatchToProps & OwnProps> = ({
  playlist,
  episodeId,
  addEpisode,
  removeEpisode,
}) => {
  const containsEpisode = playlist.members.some(
    (member) => member.episodeId === episodeId,
  )

  return (
    <div className="mb-1 px-1">
      <label className="inline-flex items-center">
        <input
          type="checkbox"
          className="form-checkbox"
          checked={containsEpisode}
          onChange={() => (containsEpisode ? removeEpisode() : addEpisode())}
        />
        <span className="ml-2">{playlist.title}</span>
      </label>
    </div>
  )
}

export default PlaylistsListItem
