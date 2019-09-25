export function formatEpisodeDuration(d: number): string {
  const [h, s] = [Math.floor(d / (60 * 60)), d % (60 * 60)]
  const m = Math.ceil(s / 60)
  return h > 0 ? `${h} hr ${m} min` : `${m} min`
}

export function formatEpisodePubDate(dateStr: string): string {
  const now = new Date()
  const pastDate = new Date(dateStr)
  const msPerMinute = 60 * 1000
  const msPerHour = msPerMinute * 60
  const msPerDay = msPerHour * 24
  const msPerWeek = msPerDay * 7
  const elapsed = +now - +pastDate

  if (elapsed < msPerMinute) {
    return `${Math.round(elapsed / 1000)} seconds ago`
  }
  if (elapsed < msPerHour) {
    return `${Math.round(elapsed / msPerMinute)} minutes ago`
  }
  if (elapsed < msPerDay) {
    return `${Math.round(elapsed / msPerHour)} hours ago`
  }
  if (elapsed < msPerWeek) {
    return `${Math.round(elapsed / msPerDay)} days ago`
  }

  if (pastDate.getFullYear() === now.getFullYear()) {
    return pastDate.toLocaleDateString('en-US', {
      day: 'numeric',
      month: 'short',
    })
  }
  return pastDate.toLocaleDateString('en-US', {
    day: 'numeric',
    month: 'short',
    year: 'numeric',
  })
}

export function formatPlayerDuration(
  current: number,
  total: number,
): [string, string] {
  const regex = total < 60 * 60 ? /\d\d:(\d\d:\d\d)/ : /(\d\d:\d\d:\d\d)/
  return [
    new Date(0, 0, 0, 0, 0, current).toTimeString().match(regex)![1],
    new Date(0, 0, 0, 0, 0, total)!.toTimeString().match(regex)![1],
  ]
}
