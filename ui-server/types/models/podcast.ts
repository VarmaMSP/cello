export type PodcastType = 'SERIAL' | 'EPISODE'

export class PodcastCategory {
  categoryId: string

  constructor(j: any) {
    this.categoryId = j['category_id']
  }
}

export class Podcast {
  id: string
  urlParam: string
  title: string
  summary: string
  description: string
  language: string
  explicit: boolean
  author: string
  totalEpisodes: number
  type: PodcastType
  complete: boolean
  earliestEpisodePubDate: string
  copyright: string
  categories: PodcastCategory[]

  constructor(j: any) {
    this.id = j['id'] || ''
    this.urlParam = j['url_param'] || ''
    this.title = j['title'] || ''
    this.summary = j['summary'] || ''
    this.description = j['description'] || ''
    this.language = j['language'] || 'en'
    this.explicit = j['explicit'] || false
    this.author = j['author'] || ''
    this.totalEpisodes = j['total_episodes'] || 0
    this.type = j['type'] || 'EPISODE'
    this.complete = j['complete'] || false
    this.earliestEpisodePubDate = j['earliest_episode_pub_date'] || ''
    this.copyright = j['copyright'] || ''
    this.categories = (j['categories'] || []).map(
      (o: any) => new PodcastCategory(o),
    )
  }
}
