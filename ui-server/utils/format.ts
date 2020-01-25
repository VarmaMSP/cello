import format from 'date-fns/format'

export function now(): string {
  return format(new Date(), 'yyyy-MM-dd HH:mm:ss')
}

export function formatVolume(value: number): string {
  return `${Math.trunc(roundTo(value, 2) * 100)}%`
}

export function formatPlaybackRate(rate: number): string {
  return `${roundTo(rate, 2)}x`
}

export function formatEpisodeDuration(d: number): string {
  const [h, s] = [Math.floor(d / (60 * 60)), d % (60 * 60)]
  const m = Math.ceil(s / 60)
  return h > 0 ? `${h} hr ${m} min` : `${m} min`
}

export function formatEpisodePubDate(
  dateStr: string,
  relative: boolean = true,
): string {
  const now = new Date()
  const pastDate = new Date(dateStr)

  if (relative) {
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

export function formatDuration(d: number): string {
  const regex = d < 60 * 60 ? /\d\d:(\d\d:\d\d)/ : /(\d\d:\d\d:\d\d)/
  return new Date(0, 0, 0, 0, 0, d)!.toTimeString().match(regex)![1]
}

function roundTo(n: number, digits: number) {
  var multiplicator = Math.pow(10, digits)
  n = parseFloat((n * multiplicator).toFixed(11))
  var test = Math.round(n) / multiplicator
  return +test.toFixed(digits)
}

export function formatCategoryTitle(str: string): string {
  return str
    .toLowerCase()
    .replace(/& /g, '')
    .replace(/ /g, '-')
}

export function getIdFromUrlParam(s: string): string {
  return s.split('-').splice(-1)[0]
}
