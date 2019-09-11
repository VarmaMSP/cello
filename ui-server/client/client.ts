import fetch from 'isomorphic-unfetch'
import { Curation, Episode, Podcast } from 'types/app'

export default class Client {
  url: string

  constructor(url: string) {
    this.url = url
  }

  getPodcastRoute(): string {
    return `${this.url}/podcasts`
  }

  getResultsRoute(): string {
    return `${this.url}/results`
  }

  getCurationsRoute(): string {
    return `${this.url}/curations`
  }

  async doFetch(method: string, url: string, body?: object): Promise<any> {
    let data: object
    let response: Response

    try {
      response = await fetch(url, {
        method,
        body: body ? JSON.stringify(body) : undefined,
        credentials: 'include',
      })
      data = await response.json()
    } catch (err) {
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

  async searchPodcasts(searchQuery: string): Promise<{ podcasts: Podcast[] }> {
    const url = `${this.getResultsRoute()}?search_query=${searchQuery}`
    const res = await this.doFetch('GET', url)
    return {
      podcasts: res.results,
    }
  }

  async getPodcastCurations(): Promise<{
    podcastCurations: { curation: Curation; podcasts: Podcast[] }[]
  }> {
    const url = `${this.getCurationsRoute()}`
    const res = await this.doFetch('GET', url)
    return {
      podcastCurations: res.results,
    }
  }
}
