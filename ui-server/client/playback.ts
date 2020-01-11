import { Playback } from 'types/app'
import * as unmarshal from 'utils/entities'
import { doFetch } from './fetch'

export async function getPlaybacks(
  episodeIds: string[],
): Promise<{ playbacks: Playback[] }> {
  const { data } = await doFetch({
    method: 'POST',
    urlPath: `/ajax/service?endpoint=get_playbacks`,
    body: { episode_ids: episodeIds },
  })

  return {
    playbacks: (data.playbacks || []).map(unmarshal.playback),
  }
}

export async function startPlayback(episodeId: string): Promise<void> {
  await doFetch({
    method: 'POST',
    urlPath: `/ajax/service?endpoint=playback_sync&action=playback_begin`,
    body: { episode_id: episodeId },
  })
}

export async function syncPlayback(
  episodeId: string,
  position: number,
): Promise<void> {
  await doFetch({
    method: 'POST',
    urlPath: `/ajax/service?endpoint=playback_sync&action=playback_progress`,
    body: { episode_id: episodeId, position: Number(position.toFixed(6)) },
  })
}
