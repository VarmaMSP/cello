export function qs(obj: { [key: string]: any }): string {
  return Object.keys(obj)
    .map((key) => `${key}=${obj[key]}`)
    .join('&')
}

export function getIdFromUrlParam(s: string): string {
  return s.split('-').splice(-1)[0]
}