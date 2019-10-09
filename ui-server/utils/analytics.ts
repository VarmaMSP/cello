import ReactGA from 'react-ga'

export function initGA(trackingId: string) {
  ReactGA.initialize(trackingId)
}

export function logPageView() {
  ReactGA.set({ page: window.location.pathname })
  ReactGA.pageview(window.location.pathname)
}

export function logEvent(category: string = '', action: string = '') {
  if (category && action) {
    ReactGA.event({ category, action })
  }
}

export function logException(description: string = '', fatal: boolean = false) {
  if (description) {
    ReactGA.exception({ description, fatal })
  }
}
