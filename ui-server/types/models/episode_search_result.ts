export class EpisodeSearchResult {
  id: string
  urlParam: string
  title: string
  description: string

  constructor(j: any) {
    this.id = j['id'] || ''
    this.urlParam = j['url_param'] || ''
    this.title = j['title'] || ''
    this.description = j['description'] || ''
  }
}
