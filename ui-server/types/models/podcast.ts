export type PodcastType = 'SERIAL' | 'EPISODE'

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

  constructor(j: any) {
    this.id = j.id || ''
    this.urlParam = j.url_param || ''
    this.title = j.title || ''
    this.summary = j.summary || ''
    this.description = j.description || ''
    this.language = j.language || 'en'
    this.explicit = j.explicit || false
    this.author = j.author || ''
    this.totalEpisodes = j.total_episodes || 0
    this.type = j.type || 'EPISODE'
    this.complete = j.complete || false
    this.earliestEpisodePubDate = j.earliest_episode_pub_date || ''
    this.copyright = j.copyright || ''
  }

  merge(p: Podcast): Podcast {
    return <Podcast>{
      id: p.id,
      urlParam: p.urlParam,
      title: p.title,
      summary: p.summary !== '' ? p.summary : this.summary,
      description: p.description !== '' ? p.description : this.description,
      language: p.language,
      explicit: p.explicit,
      author: p.author,
      totalEpisodes: p.totalEpisodes,
      type: p.type,
      complete: p.complete,
      earliestEpisodePubDate: p.earliestEpisodePubDate,
      copyright: p.copyright,
    }
  }
}
