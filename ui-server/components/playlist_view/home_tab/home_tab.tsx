import React from 'react'
import { Playlist } from 'types/models'
import PlaylistEpisodeListItem from '../playlist_episode_list_item'

export interface OwnProps {
  playlist: Playlist
}

const HomeTab: React.FC<OwnProps> = ({ playlist }) => {
  return (
    <div>
      <h2 className="font-medium tracking-wider mb-2">{'Description'}</h2>
      {!!playlist.description ? (
        <div className="tracking-wide">{playlist.description}</div>
      ) : (
        <div className="text-center text-gray-700 text-sm">
          {'No description'}
        </div>
      )}
      <hr className="my-6" />
      <h2 className="font-medium tracking-wider mb-6">{'Episodes'}</h2>
      {playlist.members.length > 0 ? (
        <div className="text-gray-800 tracking-wide">
          {playlist.members.map(({ episodeId }, i) => (
            <PlaylistEpisodeListItem
              playlistId={playlist.id}
              position={i + 1}
              episodeId={episodeId}
            />
          ))}
        </div>
      ) : (
        <div className="text-center text-gray-700 text-sm mb-10">
          {'No episodes'}
        </div>
      )}
    </div>
  )
}

export default HomeTab
