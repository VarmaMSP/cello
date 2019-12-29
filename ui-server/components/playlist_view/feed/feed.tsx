import GridResponsive from 'components/grid_responsive'
import { PlaylistLink } from 'components/link'
import React from 'react'
import { Playlist } from 'types/app'
import { getImageUrl } from 'utils/dom'

export interface StateToProps {
  playlists: Playlist[]
  receivedAll: boolean
  isLoadingMore: boolean
}

export interface DispatchToProps {
  loadMore: (offset: number) => void
}

const Feed: React.FC<StateToProps & DispatchToProps> = ({ playlists }) => {
  return (
    <div>
      <h1 className="text-xl text-gray-900">{'Your Playlists'}</h1>
      <hr className="mt-2 mb-6 border-gray-400" />
      <GridResponsive cols={{ LG: 3, MD: 3, SM: 2 }}>
        {playlists.map((p) => (
          <div key={p.id} className="flex-none mb-10 md:px-3">
            <PlaylistLink playlistUrlParam={p.urlParam}>
              <a className="block flex items-center mb-2 mx-auto cursor-pointer">
                <img
                  className="w-24 h-24 object-contain rounded-lg border"
                  src={getImageUrl(p.previewImage)}
                />
                <div className="w-2 h-20 bg-gray-400 rounded-r border-l border-white" />
                <div className="w-2 h-16 bg-gray-300 rounded-r border-l border-white" />
                <div className="ml-3">
                  <div className="text-2xl text-gray-600 text-center">
                    {p.episodeCount}
                  </div>
                  <div className="text-2xs text-gray-600">{'episodes'}</div>
                </div>
              </a>
            </PlaylistLink>
            <PlaylistLink playlistUrlParam={p.urlParam}>
              <a className="block text-sm text-gray-800 tracking-wide line-clamp-2 hover:text-gray-900 cursor-pointer">
                {p.title}
              </a>
            </PlaylistLink>
          </div>
        ))}
      </GridResponsive>
    </div>
  )
}

export default Feed
