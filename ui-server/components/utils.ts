export function humanizeDuration(d: number): string {
  const [h, s] = [Math.floor(d / (60 * 60)), d % (60 * 60)]
  const [m] = [Math.floor(s / 60), s % 60]

  let res = `${m} min`
  if (h > 0) {
    res = `${h} hr ${res}`
  }
  return res
}
