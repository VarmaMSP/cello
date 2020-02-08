import { Category, Curation, Podcast } from 'types/models'
import { formatCategoryTitle } from 'utils/format'
import { doFetch } from './fetch'

export async function getHomePageData(): Promise<{
  categories: Curation[]
  podcasts: Podcast[]
  x: Category[]
}> {
  const { raw, categories } = await doFetch({
    method: 'GET',
    urlPath: '/',
  })

  let x = <Curation[]>[]
  for (let i = 0; i < raw.categories.length; ++i) {
    let tmp = raw.categories[i]
    let cat = <Curation>{
      id: formatCategoryTitle(tmp.title),
      title: tmp.title,
      members: [],
    }
    x.push(cat)

    if (!!tmp.sub) {
      for (let j = 0; j < tmp.sub.length; ++j) {
        x.push(<Curation>{
          id: `${cat.id}-${formatCategoryTitle(tmp.sub[j])}`,
          parentId: cat.id,
          title: tmp.sub[j],
          members: [],
        })
      }
    }
  }

  return {
    categories: x,
    x: categories,
    podcasts: (raw.recommended || []).map((o: any) => new Podcast(o)),
  }
}

export async function getChartPageData(
  chartId: string,
): Promise<{
  podcasts: Podcast[]
}> {
  const { raw } = await doFetch({
    method: 'GET',
    urlPath: `/charts/${chartId}`,
  })

  return { podcasts: (raw.podcasts || []).map((o: any) => new Podcast(o)) }
}
