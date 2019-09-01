export default class Client {
  url: string = ''

  getPodcastRoute(): string {
    return `${this.url}/podcasts`
  }

  async doFetch(method: string, url: string, body?: object): Promise<any> {
    const request = new Request(url, {
      method,
      body: body ? JSON.stringify(body) : undefined,
      credentials: 'include',
    })

    let response: Response
    let data: { error: Error; message: any }
    try {
      response = await fetch(request)
      data = await response.json()
    } catch (err) {
      console.log(err)
      throw new Error(err.string())
    }

    if (data!.error) {
      if (response!.status === 401) {
        throw new Error('error')
      } else {
        throw new Error('error')
      }
    }
    return data
  }

  async getPodcastById(podcastId: string): Promise<any> {
    return this.doFetch('GET', `${this.getPodcastRoute()}/${podcastId}`)
  }
}
