export class PodcastSearchResult {
  id: string
  urlParam: string
  title: string
  author: string
  description: string

  constructor(j: any) {
    this.id = j['id'] || ''
    this.urlParam = j['url_param'] || ''
    this.title = j['title'] || ''
    this.author = j['author'] || ''
    this.description = j['description'] || ''
  }
}
