export type AudioState = 'PLAYING' | 'PAUSED' | 'LOADING' | 'ENDED'

export interface Entity {
  id: string
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
