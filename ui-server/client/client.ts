import fetch from 'isomorphic-unfetch'
import { Curation, Episode, Podcast, User } from 'types/app'
import { unmarshalEpisode, unmarshalPodcast } from 'utils/entities'

export interface RequestException {
  url: string
  statusCode: number
  responseHeaders: { [key: string]: string }
  err?: string
}

export default class Client {
  baseUrl: string

  constructor(url: string) {
    this.baseUrl = url
  }

  url = () => `${this.baseUrl}${process.browser ? '/api' : ''}`

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
    const res = await this.doFetch('GET', `${this.url()}/podcasts/${podcastId}`)
    return {
      podcast: unmarshalPodcast(res.podcast),
      episodes: (res.episodes || []).map(unmarshalEpisode),
    }
  }

  async subscribeToPodcast(podcastId: string): Promise<void> {
    await this.doFetch('PUT', `${this.url()}/podcasts/${podcastId}/subscribe`)
  }

  async unsubscribeToPodcast(podcastId: string): Promise<void> {
    await this.doFetch('PUT', `${this.url()}/podcasts/${podcastId}/unsubscribe`)
  }

  async searchPodcasts(searchQuery: string): Promise<{ podcasts: Podcast[] }> {
    const res = await this.doFetch(
      'GET',
      `${this.url()}/results?search_query=${searchQuery}`,
    )
    return {
      podcasts: (res.results || []).map(unmarshalPodcast),
    }
  }

  async getPodcastCurations(): Promise<{
    podcastCurations: { curation: Curation; podcasts: Podcast[] }[]
  }> {
    const res = await this.doFetch('GET', `${this.url()}/curations`)
    return {
      podcastCurations: (res.results || []).map((r: any) => ({
        curation: r.curation,
        podcasts: (r.podcasts || []).map(unmarshalPodcast),
      })),
    }
  }

  async getSignedInUser(): Promise<{ user: User; subscriptions: Podcast[] }> {
    const res = await this.doFetch('GET', `${this.url()}/me`)
    return {
      user: res.user,
      subscriptions: (res.subscriptions || []).map(unmarshalPodcast),
    }
  }

  async signOutUser(): Promise<void> {
    await this.doFetch('GET', `${this.url()}/signout`)
  }

  async getUserFeed(): Promise<{ episodes: Episode[] }> {
    const res = await this.doFetch('GET', `${this.url()}/feed`)
    return {
      episodes: (res.episodes || []).map(unmarshalEpisode),
    }
  }

  async getTrendingPodcasts(): Promise<Podcast[]> {
    const res = await this.doFetch(
      'GET',
      `${this.baseUrl}/static/trending.json`,
    )
    return res.map(unmarshalPodcast)
  }
}
