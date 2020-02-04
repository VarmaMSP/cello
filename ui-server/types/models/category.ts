export class Category {
  id: string
  name: string
  parentId?: string

  constructor(j: any) {
    this.id = j['id']
    this.name = j['name']
    this.parentId = j['parent_id']
  }
}
