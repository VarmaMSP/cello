import React from 'react'
import { Playlist } from 'types/app'
import PlaylistEpisodeListItem from '../playlist_episode_list_item'

export interface OwnProps {
  playlist: Playlist
}

const HomeTab: React.FC<OwnProps> = ({ playlist }) => {
  return (
    <div>
      <h2 className="font-medium tracking-wide mb-2">{'Description'}</h2>
      {!!playlist.description ? (
        <div className="tracking-wide mb-10">{playlist.description}</div>
      ) : (
        <div className="text-center text-gray-700 text-sm mb-10">
          {'No description'}
        </div>
      )}

      <h2 className="font-medium tracking-wide mb-4">{'Episodes'}</h2>
      {playlist.members.length > 0 ? (
        <div className="text-gray-800 tracking-wide">
          {playlist.members.map(({ episodeId }, i) => (
            <PlaylistEpisodeListItem position={i + 1} episodeId={episodeId} />
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
