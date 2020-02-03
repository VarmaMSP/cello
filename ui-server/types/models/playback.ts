export class Playback {
  episodeId: string
  progress: number
  lastPlayedAt: string

  constructor(j: any) {
    this.episodeId = j.episode_id || ''
    this.progress = j.progress || 0
    this.lastPlayedAt = j.last_played_at || ''
  }
}
