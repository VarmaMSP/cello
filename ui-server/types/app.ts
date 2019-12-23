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
  userId: string
  privacy: PlaylistPrivacy
}

export interface PlaylistMember {
  id: string
  episodeId: string
  playlistId: string
}

export type CurationType = 'CATEGORY' | 'NORMAL'

export interface Curation {
  id: string
  parentId?: string
  title: string
  subTitle?: string
  type: CurationType
}

export interface CurationMember {
  id: string
  curationId: string
  podcastId: string
}
