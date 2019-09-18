export type AudioState = 'PLAYING' | 'PAUSED' | 'LOADING' | 'ENDED'

export type ScreenWidth = 'SM' | 'MD' | 'LG'

export interface Entity {
  id: string
}

export interface User extends Entity {
  name: string
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
