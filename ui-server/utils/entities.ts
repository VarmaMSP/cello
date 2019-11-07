import { Episode, EpisodePlayback, Playlist, Podcast, User } from 'types/app'

export function unmarshalUser(j: any): User {
  return {
    id: j.id,
    name: j.name,
    email: j.email,
  }
}

export function unmarshalPodcast(j: any): Podcast {
  return {
    id: j.id,
    title: j.title,
    author: j.author,
    description: j.description,
    type: j.type,
    complete: j.complete || 0,
  }
}

export function unmarshalEpisode(j: any): Episode {
  return {
    id: j.id,
    podcastId: j.podcast_id,
    title: j.title,
    description: j.description,
    mediaUrl: j.media_url,
    mediaType: j.media_type,
    episode: j.episode || 0,
    season: j.season || 0,
    pubDate: j.pub_date,
    duration: j.duration || 0,
  }
}

export function unmarshalEpisodePlayback(j: any): EpisodePlayback {
  return {
    id: j.episode_id,
    episodeId: j.episode_id,
    count: j.count || 0,
    currentTime: j.current_time || 0,
  }
}

export function unmarshalPlaylist(j: any): Playlist {
  return {
    id: j.id,
    title: j.title,
    privacy: j.privacy,
    createdBy: j.created_by,
  }
}
