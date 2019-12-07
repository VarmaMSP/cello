import { Podcast, PodcastList } from 'types/app'
import * as unmarshal from 'utils/entities'
import { formatCategoryTitle } from 'utils/format'
import { doFetch } from './fetch'

export async function getDiscoverPageData(): Promise<{
  categories: PodcastList[]
  podcasts: Podcast[]
}> {
  const { data } = await doFetch({
    method: 'GET',
    urlPath: '/discover',
  })

  let categories = <PodcastList[]>[]
  for (let i = 0; i < data.categories.length; ++i) {
    let tmp = data.categories[i]
    let cat = <PodcastList>{
      id: formatCategoryTitle(tmp.title),
      title: tmp.title,
      subTitle: tmp.sub_title,
    }
    for (let j = 0; j < tmp.sub.length; ++j) {
      categories.push(<PodcastList>{
        id: `${cat.id}-${formatCategoryTitle(tmp.sub[j])}`,
        parentId: cat.id,
        title: tmp.sub[j],
      })
    }
    categories.push(cat)
  }
  
  return {
    categories,
    podcasts: (data.recommended || []).map(unmarshal.podcast),
  }
}
