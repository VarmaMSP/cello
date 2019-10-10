import { Episode, Podcast } from 'types/app'

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
