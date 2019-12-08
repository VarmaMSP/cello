export function getPodcastEpisodes(podcastId: string) {
  return `1_${podcastId}`
}

export function getSubscriptionsFeed() {
  return '2'
}

export function getHistoryFeed() {
  return '3'
}

export function getDiscoverPageData() {
  return '4'
}

export function getPodcastsInList(listId: string) {
  return `5_${listId}`
}
