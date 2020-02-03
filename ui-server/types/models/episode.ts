type EpisodeType = 'TRAILER' | 'BONUS' | 'FULL'

export class Episode {
  id: string
  urlParam: string
  podcastId: string
  title: string
  summary: string
  mediaUrl: string
  pubDate: string
  description: string
  duration: number
  explicit: boolean
  episode: number
  season: number
  type: EpisodeType
  progress: number
  lastPlayedAt: string

  constructor(j: any) {
    this.id = j['id'] || ''
    this.urlParam = j['url_param'] || ''
    this.podcastId = j['podcast_id'] || ''
    this.title = j['title'] || ''
    this.summary = j['summary'] || ''
    this.mediaUrl = j['media_url'] || ''
    this.pubDate = j['pub_date'] || ''
    this.description = j['description'] || ''
    this.duration = j['duration'] || 0
    this.explicit = j['explicit'] || false
    this.episode = j['episode'] || 0
    this.season = j['season'] || 0
    this.type = j['type'] || 'FULL'
    this.progress = j['progress'] || 0
    this.lastPlayedAt = j['last_played_at'] || ''
  }
}
