export interface Podcast {
  id: string
  description: string
  author: string
}

export interface Episode {
  id: string
  podcastId: string
  description: string
  duration: number
}
