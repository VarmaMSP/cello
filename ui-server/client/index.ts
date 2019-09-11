import Client from './client'

const client = new Client(
  !!process.browser ? 'http://localhost:8080' : 'http://localhost:8080',
)

export default client
