import fetch from 'isomorphic-unfetch'
import {
  Curation,
  Episode,
  EpisodePlayback,
  Playlist,
  Podcast,
  User,
} from 'types/app'
import {
  unmarshalEpisode,
  unmarshalEpisodePlayback,
  unmarshalPlaylist,
  unmarshalPodcast,
  unmarshalUser,
} from 'utils/entities'

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

  async getEpisode(
    episodeId: string,
  ): Promise<{ podcast: Podcast; episode: Episode }> {
    const res = await this.doFetch('GET', `${this.url()}/episodes/${episodeId}`)
    return {
      podcast: unmarshalPodcast(res.podcast),
      episode: unmarshalEpisode(res.episode),
    }
  }

  async getPodcastEpisodes(
    podcastId: string,
    limit: number,
    offset: number,
    order: 'pub_date_desc' | 'pub_date_asc',
  ): Promise<{
    episodes: Episode[]
    playbacks: EpisodePlayback[]
  }> {
    const res = await this.doFetch(
      'GET',
      `${this.url()}/podcasts/${podcastId}/episodes?limit=${limit}&offset=${offset}&order=${order}`,
    )
    return {
      episodes: (res.episodes || []).map(unmarshalEpisode),
      playbacks: (res.playbacks || []).map(unmarshalEpisodePlayback),
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

  async getSignedInUser(): Promise<{
    user: User | undefined
    subscriptions: Podcast[]
  }> {
    const res = await this.doFetch('GET', `${this.url()}/me`)
    return {
      user: res.user && unmarshalUser(res.user),
      subscriptions: (res.subscriptions || []).map(unmarshalPodcast),
    }
  }

  async signOutUser(): Promise<void> {
    await this.doFetch('GET', `${this.url()}/signout`)
  }

  async getSubscriptionsFeed(
    offset: number,
    limit: number,
  ): Promise<{
    episodes: Episode[]
    playbacks: EpisodePlayback[]
  }> {
    const res = await this.doFetch(
      'GET',
      `${this.url()}/subscriptions/feed?limit=${limit}&offset=${offset}`,
    )
    return {
      episodes: (res.episodes || []).map(unmarshalEpisode),
      playbacks: (res.playbacks || []).map(unmarshalEpisodePlayback),
    }
  }

  async getHistoryFeed(
    offset: number,
    limit: number,
  ): Promise<{
    podcasts: Podcast[]
    episodes: Episode[]
    playbacks: EpisodePlayback[]
  }> {
    const res = await this.doFetch(
      'GET',
      `${this.url()}/history/feed?limit=${limit}&offset=${offset}`,
    )
    return {
      podcasts: (res.podcasts || []).map(unmarshalPodcast),
      episodes: (res.episodes || []).map(unmarshalEpisode),
      playbacks: (res.playbacks || []).map(unmarshalEpisodePlayback),
    }
  }

  async getTrendingPodcasts(): Promise<Podcast[]> {
    const res = await this.doFetch('GET', `${this.url()}/trending`)
    return res.map(unmarshalPodcast)
  }

  async syncPlayback(episodeId: string): Promise<void> {
    await this.doFetch('POST', `${this.url()}/sync/${episodeId}`)
  }

  async syncPlaybackProgress(
    episodeId: string,
    currentTime: number,
  ): Promise<void> {
    await this.doFetch('POST', `${this.url()}/sync/${episodeId}/progress`, {
      current_time: Math.floor(currentTime),
    })
  }

  async getEpisodePlaybacks(
    episodeIds: string[],
  ): Promise<{ playbacks: EpisodePlayback[] }> {
    const res = await this.doFetch('PUT', `${this.url()}/playback`, {
      episode_ids: episodeIds,
    })
    return {
      playbacks: (res.playbacks || []).map(unmarshalEpisodePlayback),
    }
  }

  async getCurrentUserPlaylists(): Promise<{ playlists: Playlist[] }> {
    const res = await this.doFetch('GET', `${this.url()}/playlists`)
    return {
      playlists: (res.playlists || []).map(unmarshalPlaylist),
    }
  }

  async createPlaylist(
    title: string,
    privacy: string,
  ): Promise<{ playlist: Playlist }> {
    const res = await this.doFetch('POST', `${this.url()}/playlists`, {
      title,
      privacy,
    })
    return {
      playlist: unmarshalPlaylist(res.playlist),
    }
  }
}
