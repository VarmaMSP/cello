import Client from './client'

let baseUrl = 'http://localhost:8081'
if (process.env.NODE_ENV === 'development' && process.browser) {
  baseUrl = 'http://localhost:8080'
}
if (process.env.NODE_ENV === 'production' && process.browser) {
  baseUrl = 'https://phenopod.com'
}

const client = new Client(baseUrl)

export default client
