export type AudioState = 'PLAYING' | 'PAUSED' | 'LOADING' | 'ENDED'

export type ViewportSize = 'SM' | 'MD' | 'LG'

export type Modal =
  | { type: 'NONE' }
  | { type: 'SIGNIN_MODAL' }
  | { type: 'EPISODE_MODAL'; episodeId: string }
  | { type: 'ADD_TO_PLAYLIST_MODAL'; episodeId: string }

export interface Entity {
  id: string
}

export interface User extends Entity {
  name: string
  email: string
}

export interface Curation extends Entity {
  title: string
}

export interface Podcast extends Entity {
  title: string
  author: string
  description: string
  type: string
  complete: number
}

export interface Episode extends Entity {
  podcastId: string
  title: string
  description: string
  mediaUrl: string
  mediaType: string
  episode: number
  season: number
  pubDate: string
  duration: number
}

export interface EpisodePlayback extends Entity {
  episodeId: string
  count: number
  currentTime: number
}

export interface Playlist extends Entity {
  title: string
  createdBy: string
  privacy: 'PUBLIC' | 'PRIVATE' | 'ANONYMOUS'
}
