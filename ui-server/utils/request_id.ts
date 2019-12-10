export function getPodcastEpisodes(podcastId: string) {
  return `1_${podcastId}`
}

export function getSubscriptionsFeed() {
  return '2'
}

export function getHistoryFeed() {
  return '3'
}

export function getHomePageData() {
  return '4'
}

export function getPodcastsInChart(chartId: string) {
  return `5_${chartId}`
}

export function createPlaylist() {
  return '6'
}