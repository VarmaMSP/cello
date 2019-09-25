import fetch from 'isomorphic-unfetch'
import { Curation, Episode, Podcast, User } from 'types/app'
import { episodeFromJson, podcastFromJson } from 'utils/entities'

export interface RequestException {
  url: string
  statusCode: number
  responseHeaders: { [key: string]: string }
  err?: string
}

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
    // Make Request
    let response: Response
    try {
      response = await fetch(url, {
        method,
        body: body ? JSON.stringify(body) : undefined,
        credentials: 'include',
      })
    } catch (err) {
      throw <RequestException>{
        url: url,
        statusCode: 500,
        responseHeaders: {},
        err: err.toString(),
      }
    }

    // check for non 2xx { not OK }
    if (response!.status.toString()[0] !== '2') {
      throw <RequestException>{
        url: url,
        statusCode: response!.status,
        responseHeaders: {},
      }
    }

    // Parse json body
    let data: object = {}
    try {
      if (response!.headers.get('Content-Type') === 'application/json') {
        data = await response.json()
      }
    } catch (err) {
      throw <RequestException>{
        url: url,
        statusCode: response!.status,
        responseHeaders: {},
        err: 'Error Parsing JSON response',
      }
    }

    return data
  }

  async getPodcastById(
    podcastId: string,
  ): Promise<{ podcast: Podcast; episodes: Episode[] }> {
    const url = `${this.getPodcastRoute()}/${podcastId}`
    const res = await this.doFetch('GET', url)
    return {
      podcast: podcastFromJson(res.podcast),
      episodes: (res.episodes || []).map(episodeFromJson),
    }
  }

  async searchPodcasts(searchQuery: string): Promise<{ podcasts: Podcast[] }> {
    const url = `${this.getResultsRoute()}?search_query=${searchQuery}`
    const res = await this.doFetch('GET', url)
    return {
      podcasts: (res.results || []).map(podcastFromJson),
    }
  }

  async getPodcastCurations(): Promise<{
    podcastCurations: { curation: Curation; podcasts: Podcast[] }[]
  }> {
    const url = `${this.getCurationsRoute()}`
    const res = await this.doFetch('GET', url)
    return {
      podcastCurations: (res.results || []).map((r: any) => ({
        curation: r.curation,
        podcasts: (r.podcasts || []).map(podcastFromJson),
      })),
    }
  }

  async getSignedInUser(): Promise<{ user: User; subscriptions: Podcast[] }> {
    const url = `${this.url}/me`
    const res = await this.doFetch('GET', url)
    return {
      user: res.user,
      subscriptions: (res.subscriptions || []).map(podcastFromJson),
    }
  }

  async signOutUser(): Promise<void> {
    const url = `${this.url}/signout`
    await this.doFetch('GET', url)
  }

  async subscribeToPodcast(podcastId: string): Promise<void> {
    const url = `${this.getPodcastRoute()}/${podcastId}/subscribe`
    await this.doFetch('PUT', url)
  }

  async unsubscribeToPodcast(podcastId: string): Promise<void> {
    const url = `${this.getPodcastRoute()}/${podcastId}/unsubscribe`
    await this.doFetch('PUT', url)
  }

  async getTrendingPodcasts(): Promise<Podcast[]> {
    const url = `http://localhost:8080/static/trending.json`
    const res = await this.doFetch('GET', url)
    return res.map(podcastFromJson)
  }

  async getUserFeed(): Promise<{ episodes: Episode[] }> {
    const url = `${this.url}/feed`
    const res = await this.doFetch('GET', url)
    return {
      episodes: (res.episodes || []).map(episodeFromJson),
    }
  }
}
