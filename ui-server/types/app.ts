export type AudioState = 'PLAYING' | 'PAUSED' | 'LOADING' | 'ENDED'

export type ViewportSize = 'SM' | 'MD' | 'LG'

export type Modal =
  | { type: 'NONE' }
  | { type: 'SIGNIN_MODAL' }
  | { type: 'ADD_TO_PLAYLIST_MODAL'; episodeId: string }
  | { type: 'CREATE_PLAYLIST_MODAL' }

export interface Podcast {
  id: string
  urlParam: string
  title: string
  description: string
  language: string
  explicit: boolean
  author: string
  totalEpisodes: number
  type: 'SERIAL' | 'EPISODE'
  complete: number
}

export interface Episode {
  id: string
  urlParam: string
  podcastId: string
  title: string
  mediaUrl: string
  pubDate: string
  description: string
  duration: number
  explicit: boolean
  episode: number
  season: number
  type: 'TRAILER' | 'BONUS' | 'FULL'
  progress: number
  lastPlayedAt: string
}

export interface Playback {
  episodeId: string
  progress: number
  lastPlayedAt: string
}

export interface User {
  id: string
  name: string
  email: string
}

export interface Playlist {
  id: string
  title: string
  userId: string
  privacy: 'PUBLIC' | 'PRIVATE' | 'ANONYMOUS'
}
