import { Curation, Podcast } from 'types/app'
import * as unmarshal from 'utils/entities'
import { formatCategoryTitle } from 'utils/format'
import { doFetch } from './fetch'

export async function getHomePageData(): Promise<{
  categories: Curation[]
  podcasts: Podcast[]
}> {
  const { data } = await doFetch({
    method: 'GET',
    urlPath: '/',
  })

  let categories = <Curation[]>[]
  for (let i = 0; i < data.categories.length; ++i) {
    let tmp = data.categories[i]
    let cat = <Curation>{
      id: formatCategoryTitle(tmp.title),
      title: tmp.title,
      members: [],
    }
    categories.push(cat)

    if (!!tmp.sub) {
      for (let j = 0; j < tmp.sub.length; ++j) {
        categories.push(<Curation>{
          id: `${cat.id}-${formatCategoryTitle(tmp.sub[j])}`,
          parentId: cat.id,
          title: tmp.sub[j],
          members: [],
        })
      }
    }
  }

  return {
    categories,
    podcasts: (data.recommended || []).map(unmarshal.podcast),
  }
}
