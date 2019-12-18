import { Podcast } from 'types/app'
import * as unmarshal from 'utils/entities'
import { doFetch } from './fetch'

export async function getChartPageData(
  chartId: string,
): Promise<{
  podcasts: Podcast[]
}> {
  const { data } = await doFetch({
    method: 'GET',
    urlPath: `/charts/${chartId}`,
  })

  return { podcasts: (data.podcasts || []).map(unmarshal.podcast) }
}
