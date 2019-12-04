import { Playback } from 'types/app'
import * as unmarshal from 'utils/entities'
import { doFetch } from './fetch'

export async function getPlaybacks(
  episodeIds: string[],
): Promise<{ playbacks: Playback[] }> {
  const { data } = await doFetch({
    method: 'GET',
    urlPath: `/playback?ids=${episodeIds.join(',')}`,
  })

  return {
    playbacks: (data.playbacks || []).map(unmarshal.playback),
  }
}

export async function startPlayback(episodeId: string): Promise<void> {
  await doFetch({
    method: 'POST',
    urlPath: `/playback/${episodeId}`,
  })
}

export async function syncPlayback(episodeId: string): Promise<void> {
  await doFetch({
    method: 'POST',
    urlPath: `/playback/${episodeId}/sync`,
  })
}
