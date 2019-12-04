import { Episode, Playlist } from 'types/app'
import * as unmarshal from 'utils/entities'
import { doFetch } from './fetch'

export async function getUserPlaylists(): Promise<{ playlists: Playlist[] }> {
  const { data } = await doFetch({
    method: 'GET',
    urlPath: `/playlists`,
  })

  return { playlists: (data.playlists || []).map(unmarshal.playlist) }
}

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

export async function createPlaylist(
  title: string,
  privacy: string,
): Promise<{ urlParam: string }> {
  const { responseHeaders } = await doFetch({
    method: 'POST',
    urlPath: `/playlists`,
    body: { title, privacy },
  })

  return { urlParam: responseHeaders['location'] }
}

export async function addEpisodeToPlaylist(
  episodeId: string,
  playlistId: string,
): Promise<void> {
  await doFetch({
    method: 'POST',
    urlPath: `/playlists/${playlistId}/episodes/${episodeId}`,
  })
}
