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

export async function getPlaylistsPageData(): Promise<{
  playlists: Playlist[]
}> {
  const { data } = await doFetch({
    method: 'GET',
    urlPath: `/playlists`,
  })

  return { playlists: (data.playlists || []).map(unmarshal.playlist) }
}

export async function serviceAddToPlaylist(
  episodeId: string,
): Promise<{
  playlists: Playlist[]
}> {
  const { data } = await doFetch({
    method: 'POST',
    urlPath: '/ajax/service?endpoint=add_to_playlist',
    body: {
      episode_ids: [episodeId],
    },
  })

  return { playlists: (data.playlists || []).map(unmarshal.playlist) }
}

export async function serviceCreatePlaylist(
  title: string,
  privacy: string,
  episodeId: string,
): Promise<{ playlist: Playlist }> {
  const { data } = await doFetch({
    method: 'POST',
    urlPath: `/ajax/service?endpoint=create_playlist`,
    body: { title, privacy, description: '', episode_ids: [episodeId] },
  })

  return { playlist: unmarshal.playlist(data.playlist) }
}

export async function serviceAddEpisodeToPlaylist(
  playlistId: string,
  episodeId: string,
): Promise<void> {
  await doFetch({
    method: 'POST',
    urlPath: '/ajax/service?endpoint=edit_playlist&action=add_episode',
    body: {
      episode_id: episodeId,
      playlist_id: playlistId,
    },
  })
}

export async function serviceRemoveEpisodeFromPlaylist(
  playlistId: string,
  episodeId: string,
): Promise<void> {
  await doFetch({
    method: 'POST',
    urlPath: `/ajax/service?endpoint=edit_playlist&action=remove_episode`,
    body: {
      episode_id: episodeId,
      playlist_id: playlistId,
    },
  })
}
