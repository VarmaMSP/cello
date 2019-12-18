export function qs(obj: { [key: string]: any }): string {
  return Object.keys(obj)
    .map((key) => `${key}=${obj[key]}`)
    .join('&')
}
