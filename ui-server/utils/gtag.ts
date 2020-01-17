interface GtagEvent {
  action: string
  category: 'podcast' | 'episode' | 'playlist'
  label: string
  value?: number
}

export const GA_TRACKING_ID = 'UA-149726196-1'

export function userId(id: string) {
  ;(window as any).gtag('config', GA_TRACKING_ID, {
    user_id: id,
  })
}

export function pageview(url: string) {
  ;(window as any).gtag('config', GA_TRACKING_ID, {
    page_path: url,
  })
}

export function search(query: string) {
  ;(window as any).gtag('event', 'search', {
    search_term: query,
  })
}

export function subscribePodcast(podcastTitle: string) {
  event({
    action: 'subscribe',
    category: 'podcast',
    label: podcastTitle,
  })
}

export function playEpisode(episodeTitle: string, startTime: number) {
  event({
    action: 'play',
    category: 'episode',
    label: episodeTitle,
    value: startTime,
  })
}

export function createPlaylist(playlistTitle: string) {
  event({
    action: 'create',
    category: 'playlist',
    label: playlistTitle,
  })
}

export function addEpisodeToPlaylist(playlistTitle: string) {
  event({
    action: 'add_to_playlist',
    category: 'playlist',
    label: playlistTitle,
  })
}

export function event({ action, category, label, value }: GtagEvent) {
  ;(window as any).gtag('event', action, {
    event_category: category,
    event_label: label,
    value: value,
  })
}
