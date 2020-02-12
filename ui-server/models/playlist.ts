export type PlaylistPrivacy = 'PUBLIC' | 'PRIVATE' | 'UNLISTED' | 'ANONYMOUS'

export class PlaylistMember {
  episodeId: string
  position: number

  constructor(j: any) {
    this.episodeId = j['episode_id'] || ''
    this.position = j['position'] || 0
  }
}

export class Playlist {
  id: string
  urlParam: string
  title: string
  description: string
  userId: string
  privacy: PlaylistPrivacy
  updatedAt: string
  previewImage: string
  episodeCount: number
  members: PlaylistMember[]

  constructor(j: any) {
    this.id = j.id || ''
    this.urlParam = j.url_param || ''
    this.title = j.title || ''
    this.description = j.description || ''
    this.privacy = j.privacy || 'PRIVATE'
    this.previewImage = j.preview_image || ''
    this.episodeCount = j.episode_count || 0
    this.userId = j.user_id || ''
    this.updatedAt = j.updated_at || ''
    this.members = (j.members || []).map((o: any) => new PlaylistMember(o))
  }
}
