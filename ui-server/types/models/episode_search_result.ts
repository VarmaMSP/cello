export class EpisodeSearchResult {
  id: string
  title: string
  description: string

  constructor(j: any) {
    this.id = j.id || ''
    this.title = j.title || ''
    this.description = j.description || ''
  }
}
