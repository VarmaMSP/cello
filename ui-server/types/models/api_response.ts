import * as Parse from 'utils/entities'
import { Episode, EpisodeSearchResult, Playback, Playlist, Podcast, PodcastSearchResult, User } from '../app'

export class ApiResponse {
  users: User[]
  podcasts: Podcast[]
  episodes: Episode[]
  playbacks: Playback[]
  playlists: Playlist[]
  podcastSearchResults: PodcastSearchResult[]
  episodeSearchResults: EpisodeSearchResult[]

  constructor(j: any) {
    const data = (j['data'] || {}) as any

    this.users = (data['users'] || []).map(Parse.user)
    this.podcasts = (data['podcasts'] || []).map(Parse.podcast)
    this.episodes = (data['episodes'] || []).map(Parse.episode)
    this.playbacks = (data['playbacks'] || []).map(Parse.playback)
    this.playlists = (data['playlists'] || []).map(Parse.playlist)
    this.podcastSearchResults = (data['podcast_search_results'] || []).map(
      Parse.podcastSearchResult,
    )
    this.episodeSearchResults = (data['episode_search_results'] || []).map(
      Parse.episodeSearchResult,
    )
  }
}
