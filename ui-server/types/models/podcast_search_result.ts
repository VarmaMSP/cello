export class PodcastSearchResult {
  id: string
  title: string
  author: string
  description: string

  constructor(j: any) {
    this.id = j['id']
    this.title = j['title'] || ''
    this.author = j['author'] || ''
    this.description = j['description'] || ''
  }
}
