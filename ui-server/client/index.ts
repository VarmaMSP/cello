import Client from './client'

const client = new Client(process.env.API_URL as string)

export default client
