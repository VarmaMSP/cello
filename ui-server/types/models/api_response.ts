import { Episode } from './episode'
import { EpisodeSearchResult } from './episode_search_result'
import { Playback } from './playback'
import { Playlist } from './playlist'
import { Podcast } from './podcast'
import { PodcastSearchResult } from './podcast_search_result'
import { User } from './user'

export class ApiResponse {
  users: User[]
  podcasts: Podcast[]
  episodes: Episode[]
  playbacks: Playback[]
  playlists: Playlist[]
  podcastSearchResults: PodcastSearchResult[]
  episodeSearchResults: EpisodeSearchResult[]
  raw: any

  constructor(j: any) {
    const data = (j['data'] || {}) as any

    this.users = (data['users'] || []).map((o: any) => new User(o))
    this.podcasts = (data['podcasts'] || []).map((o: any) => new Podcast(o))
    this.episodes = (data['episodes'] || []).map((o: any) => new Episode(o))
    this.playbacks = (data['playbacks'] || []).map((o: any) => new Playback(o))
    this.playlists = (data['playlists'] || []).map((o: any) => new Playlist(o))
    this.podcastSearchResults = (data['podcast_search_results'] || []).map(
      (o: any) => new PodcastSearchResult(o),
    )
    this.episodeSearchResults = (data['episode_search_results'] || []).map(
      (o: any) => new EpisodeSearchResult(o),
    )
    this.raw = j['raw'] || {}
  }
}
