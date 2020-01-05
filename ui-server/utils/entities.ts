import {
  Episode,
  EpisodeSearchResult,
  Playback,
  Playlist,
  Podcast,
  PodcastSearchResult,
  User,
} from 'types/app'

export function user(j: any): User {
  return {
    id: j.id,
    name: j.name,
    email: j.email,
  }
}

export function podcast(j: any): Podcast {
  return {
    id: j.id,
    urlParam: j.url_param || '',
    title: j.title || '',
    summary: j.summary || '',
    description: j.description || '',
    language: j.language || 'en',
    explicit: j.explicit || false,
    author: j.author || '',
    totalEpisodes: j.total_episodes || 0,
    type: j.type || 'EPISODE',
    complete: j.complete || false,
    earliestEpisodePubDate: j.earliest_episode_pub_date,
    copyright: j.copyright || '',
    titleHighlighted: j.title_highlighted || '',
    authorHighlighted: j.author_highlighted || '',
    summaryHighlighted: j.summary_highlighted || '',
  }
}

export function episode(j: any): Episode {
  return {
    id: j.id,
    urlParam: j.url_param || '',
    podcastId: j.podcast_id,
    title: j.title,
    summary: j.summary || '',
    mediaUrl: j.media_url,
    pubDate: j.pub_date,
    description: j.description || '',
    duration: j.duration || 0,
    explicit: j.explicit || false,
    episode: j.episode || 0,
    season: j.season || 0,
    type: j.type || 'FULL',
    progress: j.progress || 0,
    lastPlayedAt: j.last_played_at || '',
  }
}

export function playback(j: any): Playback {
  return {
    episodeId: j.episode_id,
    progress: j.progress,
    lastPlayedAt: j.last_played_at,
  }
}

export function playlist(j: any): Playlist {
  return {
    id: j.id,
    urlParam: j.url_param || '',
    title: j.title,
    description: j.description || '',
    privacy: j.privacy || 'PRIVATE',
    previewImage: j.preview_image || '',
    episodeCount: j.episode_count || 0,
    userId: j.user_id,
    updatedAt: j.updated_at || '',
    members: (j.members || []).map((k: any) => ({
      episodeId: k.episode_id,
      position: k.position,
    })),
  }
}

export function podcastSearchResult(j: any): PodcastSearchResult {
  return {
    id: j.id,
    title: j.title || '',
    author: j.author || '',
    summary: j.summary || '',
  }
}

export function episodeSearchResult(j: any): EpisodeSearchResult {
  return {
    id: j.id,
    title: j.title || '',
    summary: j.summary || '',
  }
}
