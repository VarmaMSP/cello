import { Podcast, Episode } from 'types/app'

export default class Client {
  url: string = 'http://localhost:8080'

  getPodcastRoute(): string {
    return `${this.url}/podcasts`
  }

  async doFetch(method: string, url: string, body?: object): Promise<any> {
    const request = new Request(url, {
      method,
      body: body ? JSON.stringify(body) : undefined,
      credentials: 'include',
    })

    let response: Response | null = null
    let data: object | null = null
    try {
      response = await fetch(request)
      data = await response.json()
    } catch (err) {
      console.log(err)
      throw new Error(err.toString())
    }

    if (response!.status === 401) {
      throw new Error('error')
    }

    return data
  }

  async getPodcastById(
    podcastId: string,
  ): Promise<{ podcast: Podcast; episodes: Episode[] }> {
    const url = `${this.getPodcastRoute()}/${podcastId}`
    const res = await this.doFetch('GET', url)
    return {
      podcast: res.podcast,
      episodes: res.episodes.map((e: object) => ({
        ...e,
        podcastId: res.podcast.id,
      })),
    }
  }
}
