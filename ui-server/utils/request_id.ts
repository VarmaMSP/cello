export function getPodcastEpisodes(podcastId: string) {
  return `0_${podcastId}`
}

export function getHomePageData() {
  return '1'
}

export function getSubscriptionsPageData() {
  return '2'
}

export function getSubscriptionsFeed() {
  return '2-0'
}

export function getHistoryPageData() {
  return '3'
}

export function getHistoryFeed() {
  return '3-0'
}

export function getPlaylistPageData() {
  return '4'
}

export function getPlaylistFeed() {
  return '4-0'
}

export function getPodcastsInChart(chartId: string) {
  return `5_${chartId}`
}

export function createPlaylist() {
  return '6'
}

export function getResults() {
  return '7'
}
