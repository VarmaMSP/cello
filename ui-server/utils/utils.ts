export function qs(obj: { [key: string]: any }): string {
  return Object.keys(obj)
    .map((key) => `${key}=${obj[key]}`)
    .join('&')
}

export function getIdFromUrlParam(s: string): string {
  return s.split('-').splice(-1)[0]
}

export function mergeString(
  s1: string | undefined,
  s2: string | undefined,
): string {
  if (!!s2 && s2 !== '') {
    return s2
  }
  if (!!s1 && s1 !== '') {
    return s1
  }
  return ''
}

export function mergeNumber(
  n1: number | undefined,
  n2: number | undefined,
): number {
  if (!!n2 && n2 !== 0) {
    return n2
  }
  if (!!n1 && n1 !== 0) {
    return n1
  }
  return 0
}

export function mergeArray<T>(a1: T[] | undefined, a2: T[] | undefined): T[] {
  if (!!a2 && a2.length > 0) {
    return a2
  }
  if (!!a1 && a1.length > 0) {
    return a1
  }
  return []
}

export function uniqueId() {
  return (
    '_' +
    Math.random()
      .toString(36)
      .substr(2, 9)
  )
}

export function encodeQueryParam(s: string): string {
  return s
    .trim()
    .replace(/ /g, '+')
    .concat(s[s.length - 1] === ' ' ? '+' : '')
}
