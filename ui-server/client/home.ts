import { Chart, Podcast } from 'types/app'
import * as unmarshal from 'utils/entities'
import { formatCategoryTitle } from 'utils/format'
import { doFetch } from './fetch'

export async function getHomePageData(): Promise<{
  categories: Chart[]
  podcasts: Podcast[]
}> {
  const { data } = await doFetch({
    method: 'GET',
    urlPath: '/',
  })

  let categories = <Chart[]>[]
  for (let i = 0; i < data.categories.length; ++i) {
    let tmp = data.categories[i]
    let cat = <Chart>{
      id: formatCategoryTitle(tmp.title),
      title: tmp.title,
      subTitle: tmp.sub_title,
      type: 'CATEGORY',
    }
    for (let j = 0; j < tmp.sub.length; ++j) {
      categories.push(<Chart>{
        id: `${cat.id}-${formatCategoryTitle(tmp.sub[j])}`,
        parentId: cat.id,
        title: tmp.sub[j],
        type: 'CATEGORY',
      })
    }
    categories.push(cat)
  }

  return {
    categories,
    podcasts: (data.recommended || []).map(unmarshal.podcast),
  }
}
