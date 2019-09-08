export function humanizeDuration(d: number): string {
  const [h, s] = [Math.floor(d / (60 * 60)), d % (60 * 60)]
  const [m] = [Math.floor(s / 60), s % 60]

  let res = `${m} min`
  if (h > 0) {
    res = `${h} hr ${res}`
  }
  return res
}

export function humanizePastDate(dateStr: string): string {
  const now = new Date()
  const pastDate = new Date(dateStr)

  // return relative time if dateStr is older that i month
  const msPerMinute = 60 * 1000
  const msPerHour = msPerMinute * 60
  const msPerDay = msPerHour * 24
  const msPerMonth = msPerDay * 30
  const elapsed = +now - +pastDate

  if (elapsed < msPerMinute) {
    return `${Math.round(elapsed / 1000)} seconds ago`
  } else if (elapsed < msPerHour) {
    return `${Math.round(elapsed / msPerMinute)} minutes ago`
  } else if (elapsed < msPerDay) {
    return `${Math.round(elapsed / msPerHour)} hours ago`
  } else if (elapsed < msPerMonth) {
    return `${Math.round(elapsed / msPerDay)} days ago`
  }

  if (pastDate.getFullYear() === now.getFullYear()) {
    return pastDate.toLocaleDateString('en-US', {
      day: 'numeric',
      month: 'short',
    })
  } else {
    return pastDate.toLocaleDateString('en-US', {
      day: 'numeric',
      month: 'short',
      year: 'numeric',
    })
  }
}
