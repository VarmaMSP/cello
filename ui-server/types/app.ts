export interface Podcast {
  id: string
  title: string
  author: string
  description: string
  type: string
  complete: number
}

export interface Episode {
  id: string
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
