import { Episode, Playback, Playlist, Podcast, User } from 'types/app'

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
    description: j.description || '',
    language: j.language || 'en',
    explicit: j.explicit || false,
    author: j.author || '',
    totalEpisodes: j.total_episodes || 0,
    type: j.type || 'EPISODE',
    complete: j.complete || false,
    earliestEpisodePubDate: j.earliest_episode_pub_date,
  }
}

export function episode(j: any): Episode {
  return {
    id: j.id,
    urlParam: j.url_param || '',
    podcastId: j.podcast_id,
    title: j.title,
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
    title: j.title,
    privacy: j.privacy,
    userId: j.user_id,
  }
}
