import { Episode, Playlist } from 'types/app'
import * as unmarshal from 'utils/entities'
import { doFetch } from './fetch'

export async function getPlaylist(
  playlistId: string,
): Promise<{ playlist: Playlist; episodes: Episode[] }> {
  const { data } = await doFetch({
    method: 'GET',
    urlPath: `/playlists/${playlistId}`,
  })

  return {
    playlist: unmarshal.playlist(data.playlist),
    episodes: (data.episodes || []).map(unmarshal.episode),
  }
}

export async function getUserPlaylists(
  userId: string,
  offset: number,
  limit: number,
  order: 'create_date_desc',
): Promise<{
  playlists: Playlist[]
  episodesByPlaylist: { [playlist: string]: Episode[] }
}> {
  const { data } = await doFetch({
    method: 'GET',
    urlPath: `/playlists?full_details=true&user_id=${userId}&limit=${limit}&offset=${offset}&order=${order}`,
  })

  return {
    playlists: (data.playlists || []).map(unmarshal.playlist),
    episodesByPlaylist: Object.keys(data.episodes_by_playlist || {}).reduce<{
      [playlistId: string]: Episode[]
    }>(
      (acc, playlistId) => ({
        ...acc,
        [playlistId]: (data.episodes_by_playlist || {})[playlistId].map(
          unmarshal.episode,
        ),
      }),
      {},
    ),
  }
}

export async function getSignedInUserPlaylists(): Promise<{
  playlists: Playlist[]
}> {
  const { data } = await doFetch({
    method: 'GET',
    urlPath: `/playlists?full_details=false`,
  })

  return { playlists: (data.playlists || []).map(unmarshal.playlist) }
}

export async function createPlaylist(
  title: string,
  privacy: string,
): Promise<{ urlParam: string }> {
  const { responseHeaders } = await doFetch({
    method: 'POST',
    urlPath: `/playlists`,
    body: { title, privacy },
  })

  return { urlParam: responseHeaders['Location'] }
}

export async function addEpisodeToPlaylists(
  episodeId: string,
  playlistIds: string[],
): Promise<void> {
  await doFetch({
    method: 'POST',
    urlPath: `/playlists/episodes`,
    body: {
      episode_id: episodeId,
      playlist_ids: playlistIds,
    },
  })
}
