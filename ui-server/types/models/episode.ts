type EpisodeType = 'TRAILER' | 'BONUS' | 'FULL'
import { mergeNumber, mergeString } from 'utils/utils'

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

  static merge(e1: Episode, e2: Episode): Episode {
    return {
      id: e2.id,
      urlParam: mergeString(e1.urlParam, e2.urlParam),
      podcastId: mergeString(e1.podcastId, e2.podcastId),
      title: mergeString(e1.title, e2.title),
      summary: mergeString(e1.summary, e2.summary),
      mediaUrl: mergeString(e1.mediaUrl, e2.mediaUrl),
      pubDate: mergeString(e1.pubDate, e2.pubDate),
      description: mergeString(e1.description, e2.description),
      duration: mergeNumber(e1.duration, e2.duration),
      explicit: e1.explicit || e2.explicit,
      episode: mergeNumber(e1.episode, e2.episode),
      season: mergeNumber(e1.season, e2.season),
      type: e2.type,
      progress:
        (!!e2.lastPlayedAt && e2.lastPlayedAt !== ''
          ? e2.progress
          : e1.progress) || 0,
      lastPlayedAt: mergeString(e1.lastPlayedAt, e2.lastPlayedAt),
    }
  }

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
