interface GtagEvent {
  action: string
  category: string
  label: string
  value: number
}

export const GA_TRACKING_ID = ''

export function pageview(url: string) {
  ;(window as any).gtag('config', GA_TRACKING_ID, {
    page_path: url,
  })
}

export function event({ action, category, label, value }: GtagEvent) {
  ;(window as any).gtag('event', action, {
    event_category: category,
    event_label: label,
    value: value,
  })
}
