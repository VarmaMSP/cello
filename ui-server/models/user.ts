export class User {
  id: string
  name: string
  email: string

  constructor(j: any) {
    this.id = j.id || ''
    this.name = j.name || ''
    this.email = j.email || ''
  }
}
