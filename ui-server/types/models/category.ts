export class Category {
  id: string
  urlParam: string
  name: string
  parentId?: string
  podcastIds: string[]

  constructor(j: any) {
    this.id = j['id'] || ''
    this.urlParam = j['url_param'] || ''
    this.name = j['name'] || ''
    this.parentId = j['parent_id']
    this.podcastIds = j['podcast_ids'] || []
  }
}
