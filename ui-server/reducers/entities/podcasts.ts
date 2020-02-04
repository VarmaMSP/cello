import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { Podcast } from 'types/models'

const merge = (p1: Podcast, p2: Podcast): Podcast => ({
  id: p2.id,
  urlParam: p2.urlParam,
  title: p2.title,
  summary: !!p2.summary ? p2.summary : p1.summary,
  description: !!p2.description ? p2.description : p1.description,
  language: !!p2.language ? p2.language : p1.language,
  explicit: p2.explicit,
  author: p2.author,
  totalEpisodes: p2.totalEpisodes,
  type: p2.type,
  complete: p2.complete,
  earliestEpisodePubDate: p2.earliestEpisodePubDate,
  copyright: !!p2.copyright ? p2.copyright : p1.copyright,
  categories: p2.categories.length > 0 ? p2.categories : p1.categories,
})

const byId: Reducer<{ [podcastId: string]: Podcast }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.PODCAST_ADD:
      return action.podcasts.reduce<{
        [podcastId: string]: Podcast
      }>((acc, p) => ({ ...acc, [p.id]: merge(acc[p.id] || {}, p) }), state)

    default:
      return state
  }
}

export default combineReducers({
  byId,
})
