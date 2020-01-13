import GridResponsive from 'components/grid_responsive'
import { PlaylistLink } from 'components/link'
import format from 'date-fns/format'
import parseISO from 'date-fns/parseISO'
import React from 'react'
import { Playlist } from 'types/app'
import { getImageUrl } from 'utils/dom'

export interface StateToProps {
  playlists: Playlist[]
}

const PlaylistLibrary: React.FC<StateToProps> = ({ playlists }) => {
  if (playlists.length === 0) {
    return (
      <div className="mt-10">
        <div className="text-2xl text-gray-900 tracking-wide">
          {"You have'nt created any playlists yet."}
        </div>
        <div className="mt-2 text-default text-gray-800 tracking-wide">
          {'Use + button to create one.'}
        </div>
      </div>
    )
  }

  return (
    <div>
      <h1 className="text-xl text-gray-900">{'Your Playlists'}</h1>
      <hr className="mt-2 mb-6 border-gray-400" />
      <GridResponsive cols={{ LG: 3, MD: 3, SM: 2 }}>
        {playlists.map((p) => (
          <div key={p.id} className="flex-none mb-12 px-3">
            <PlaylistLink playlistUrlParam={p.urlParam}>
              <a className="block flex items-center mb-3 mx-auto cursor-pointer">
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
              <a className="block text-sm text-gray-800 font-medium tracking-wide line-clamp-1 hover:text-gray-900 cursor-pointer">
                {p.title}
              </a>
            </PlaylistLink>
            <div className="text-2xs text-gray-700 leading-relaxed">
              {`updated on ${formatUpdateDate(p.updatedAt)}`}
            </div>
            <span className="inline-block px-2 text-2xs text-gray-700 bg-gray-300 rounded-full">
              {p.privacy.toLowerCase()}
            </span>
          </div>
        ))}
      </GridResponsive>
    </div>
  )
}

function formatUpdateDate(updateDate: string) {
  let pubDate: string | undefined
  try {
    pubDate = format(parseISO(`${updateDate} +0000`), 'PP')
  } catch (err) {}

  return pubDate
}

export default PlaylistLibrary
