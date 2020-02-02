export type AudioState = 'PLAYING' | 'PAUSED' | 'LOADING' | 'ENDED'

export type ViewportSize = 'SM' | 'MD' | 'LG'

export type Modal =
  | { type: 'NONE' }
  | { type: 'SIGNIN_MODAL' }
  | { type: 'ADD_TO_PLAYLIST_MODAL'; episodeId: string }
  | { type: 'CREATE_PLAYLIST_MODAL'; episodeId: string }

export type PodcastType = 'SERIAL' | 'EPISODE'

export interface Podcast {
  id: string
  urlParam: string
  title: string
  summary: string
  description: string
  language: string
  explicit: boolean
  author: string
  totalEpisodes: number
  type: PodcastType
  complete: boolean
  earliestEpisodePubDate: string
  copyright: string
  titleHighlighted: string
  authorHighlighted: string
  summaryHighlighted: string
}

type EpisodeType = 'TRAILER' | 'BONUS' | 'FULL'

export interface Episode {
  id: string
  urlParam: string
  podcastId: string
  title: string
  summary: string
  mediaUrl: string
  pubDate: string
  description: string
  duration: number
  explicit: boolean
  episode: number
  season: number
  type: EpisodeType
  progress: number
  lastPlayedAt: string
}

export interface Playback {
  episodeId: string
  progress: number
  lastPlayedAt: string
}

export interface User {
  id: string
  name: string
  email: string
}

export type PlaylistPrivacy = 'PUBLIC' | 'PRIVATE' | 'UNLISTED' | 'ANONYMOUS'

export interface Playlist {
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
}

export interface PlaylistMember {
  episodeId: string
  position: number
}

export interface Curation {
  id: string
  parentId?: string
  title: string
  members: string[]
}

export interface PodcastSearchResult {
  id: string
  title: string
  author: string
  description: string
}

export interface EpisodeSearchResult {
  id: string
  title: string
  description: string
}
