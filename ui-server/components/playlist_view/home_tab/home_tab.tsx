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
      <div className="text-gray-700 tracking-wide mb-10">
        {!!playlist.description ? (
          playlist.description
        ) : (
          <div className="text-center">{'No description'}</div>
        )}
      </div>

      <h2 className="font-medium tracking-wide mb-2">{'Episodes'}</h2>
      <div className="text-gray-800 tracking-wide">
        {playlist.members.map(({ episodeId }, i) => (
          <PlaylistEpisodeListItem position={i + 1} episodeId={episodeId} />
        ))}
      </div>
    </div>
  )
}

export default HomeTab
