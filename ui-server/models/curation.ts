export class Curation {
  id: string
  title: string
  podcastIds: string[]

  constructor(j: any) {
    this.id = j['id'] || ''
    this.title = j['title'] || ''
    this.podcastIds = j['podcast_ids'] || []
  }
}
