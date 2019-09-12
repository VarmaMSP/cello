import fetch from 'isomorphic-unfetch'
import { Curation, Episode, Podcast } from 'types/app'

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

    // Check for 3** {redirections}
    if (response!.status.toString()[0] === '3') {
      throw <RequestException>{
        url: url,
        statusCode: response!.status,
        responseHeaders: { location: response!.headers.get('Location') },
      }
    }

    // Check for 4xx {user errors}
    if (response!.status.toString()[0] === '4') {
      throw <RequestException>{
        url: url,
        statusCode: response!.status,
        responseHeaders: {},
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

    // Parse body is content-type is json
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
      podcast: res.podcast,
      episodes: (res.episodes || []).map((episode: object) => ({
        ...episode,
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
